package sets

import (
	"fmt"
	"reflect"
	"sync"
)

var SetVal = struct{}{}

type Set interface {
	Values() []interface{}
	RType() reflect.Type
	String() string

	Has(item interface{}) bool
	Add(items ...interface{}) error
	Remove(items ...interface{}) error
	Equals(s2 Set) bool
	Contains(s2 Set) bool
}

type set struct {
	Set
	store map[interface{}]struct{}
	mutex sync.RWMutex
	rtype reflect.Type
}

func New(items ...interface{}) (Set, error) {
	s := set{
		store: make(map[interface{}]struct{}),
		mutex: sync.RWMutex{},
	}

	err := s.Add(items...)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (s *set) Values() []interface{} {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	list := make([]interface{}, 0, len(s.store))
	for item := range s.store {
		list = append(list, item)
	}

	return list
}

func (s *set) RType() reflect.Type {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.rtype
}

func (s *set) String() string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if s.rtype == nil {
		return fmt.Sprintf("set%v", s.Values())
	}
	return fmt.Sprintf("set<%v>%v", s.rtype.String(), s.Values())
}

func (s *set) Has(item interface{}) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if len(s.store) == 0 {
		return false
	}

	_, has := s.store[item]
	return has
}

func (s *set) Add(items ...interface{}) error {
	if len(items) == 0 {
		return nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	for idx, item := range items {
		if idx == 0 && len(s.store) == 0 {
			s.rtype = reflect.TypeOf(item)
		} else if s.rtype != reflect.TypeOf(item) {
			return fmt.Errorf("cannot add item of incorrect type: %s", reflect.TypeOf(item).String())
		}
		s.store[item] = SetVal
	}

	return nil
}

func (s *set) Remove(items ...interface{}) (err error) {
	if len(items) == 0 {
		return err
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, item := range items {
		if s.rtype != reflect.TypeOf(item) {
			err = fmt.Errorf("cannot remove item of incorrect type: %s", reflect.TypeOf(item).String())
		}
		delete(s.store, item)
	}

	return err
}

func (s *set) Equals(s2 Set) bool {
	if s2 == nil {
		return false
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if len(s.store) != len(s2.Values()) {
		return false
	}

	if s.rtype != s2.RType() {
		return false
	}

	for item := range s.store {
		if !s2.Has(item) {
			return false
		}
	}

	return true
}

func (s *set) Contains(s2 Set) bool {
	// Treat nil as empty set
	if s2 == nil {
		return true
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Every set contains the empty set (regardless of type)
	if len(s2.Values()) == 0 {
		return true
	}

	if s.rtype != s2.RType() {
		return false
	}

	if len(s2.Values()) > len(s.store) {
		return false
	}

	for _, item := range s2.Values() {
		if !s2.Has(item) {
			return false
		}
	}

	return true
}
