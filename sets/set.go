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
	Add(items ...interface{}) Set
	Remove(items ...interface{}) Set
	Copy() Set

	// package private
	readLock()
	readUnlock()
}

type set struct {
	Set
	store map[interface{}]struct{}
	mutex sync.RWMutex
	rtype reflect.Type
}

func New(items ...interface{}) Set {
	s := &set{
		store: make(map[interface{}]struct{}),
		mutex: sync.RWMutex{},
	}

	return s.Add(items...)
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

func (s *set) Add(items ...interface{}) Set {
	if len(items) == 0 {
		return s
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	for idx, item := range items {
		if idx == 0 && len(s.store) == 0 {
			s.rtype = reflect.TypeOf(item)
		} else if s.rtype != reflect.TypeOf(item) {
			panic(fmt.Errorf("cannot add item of incorrect type: %s", reflect.TypeOf(item).String()))
		}
		s.store[item] = SetVal
	}

	return s
}

func (s *set) Remove(items ...interface{}) Set {
	if len(items) == 0 {
		return s
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, item := range items {
		if s.rtype != reflect.TypeOf(item) {
			panic(fmt.Errorf("cannot remove item of incorrect type: %s", reflect.TypeOf(item).String()))
		}
		delete(s.store, item)
	}

	return s
}

// This copy is shallow. It does not deep-copy values recursively
func (s *set) Copy() Set {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return New(s.Values()...)
}

func (s *set) readLock() {
	s.mutex.RLock()
}

func (s *set) readUnlock() {
	s.mutex.RUnlock()
}
