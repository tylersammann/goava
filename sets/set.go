package sets

import (
	"fmt"
	"reflect"
	"sync"
)

var SetVal = struct{}{}

type Set interface {
	Size() int
	Values() []interface{}
	String() string

	Has(item interface{}) bool
	Add(items ...interface{}) Set
	Remove(items ...interface{}) Set
	Copy() Set

	ForEach(fn func(interface{}))
	FindFirst(fn func(interface{}) bool) interface{}

	Equals(s2 Set) bool
	Contains(s2 Set) bool
	Difference(s2 Set) Set
	Intersection(s2 Set) Set
	Union(s2 Set) Set

	// package private
	rType() reflect.Type
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

func (s *set) Size() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return len(s.store)
}

// Values method copies all values to a new slice
func (s *set) Values() []interface{} {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	list := make([]interface{}, 0, len(s.store))
	for item := range s.store {
		list = append(list, item)
	}

	return list
}

func (s *set) String() string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if s.rtype == nil {
		return fmt.Sprintf("set<>%v", s.Values())
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

	for _, item := range items {
		if s.rtype == nil {
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

	if len(s.store) == 0 {
		return s
	}

	for _, item := range items {
		if s.rtype != reflect.TypeOf(item) {
			panic(fmt.Errorf("cannot remove item of incorrect type: %s", reflect.TypeOf(item).String()))
		}
		delete(s.store, item)
	}

	// An empty set has no type. Set rtype to nil if store has 0 length
	if len(s.store) == 0 {
		s.rtype = nil
	}

	return s
}

// This copy is shallow. It does not deep-copy values recursively
func (s *set) Copy() Set {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	copySet := New()
	for item := range s.store {
		copySet.Add(item)
	}

	return copySet
}

func (s *set) ForEach(fn func(interface{})) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for item := range s.store {
		fn(item)
	}
}

// Returns first element for which callback fn returns true. Else returns nil
func (s *set) FindFirst(fn func(interface{}) bool) interface{} {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for item := range s.store {
		if fn(item) {
			return item
		}
	}
	return nil
}

// Equals returns true if s has the same elements as s2
// Does not respect different memory addresses (for that, use ==)
func (s *set) Equals(s2 Set) bool {
	if s2 == nil {
		return false
	}

	s.readLock()
	defer s.readUnlock()
	s2.readLock()
	defer s2.readUnlock()

	if s.Size() != s2.Size() {
		return false
	}

	if s.rType() != s2.rType() {
		return false
	}

	differentItem := s.FindFirst(func(item interface{}) bool {
		return !s2.Has(item)
	})

	return differentItem == nil
}

func (s *set) Contains(s2 Set) bool {
	if s2 == nil {
		return true
	}

	s.readLock()
	defer s.readUnlock()
	s2.readLock()
	defer s2.readUnlock()

	// Every set contains the empty set (regardless of type)
	if s2.Size() == 0 {
		return true
	}

	if s.rType() != s2.rType() {
		return false
	}

	if s2.Size() > s2.Size() {
		return false
	}

	differentItem := s2.FindFirst(func(item interface{}) bool {
		return !s.Has(item)
	})

	return differentItem == nil
}

func (s *set) Difference(s2 Set) Set {
	s.readLock()
	defer s.readUnlock()
	s2.readLock()
	defer s2.readUnlock()

	diff := s.Copy()
	if s2 == nil {
		return diff
	}

	s2.ForEach(func(item interface{}) {
		diff.Remove(item)
	})

	return diff
}

func (s *set) Intersection(s2 Set) Set {
	inter := New()

	if s2 == nil {
		return inter
	}

	s.readLock()
	defer s.readUnlock()
	s2.readLock()
	defer s2.readUnlock()

	s.ForEach(func(item interface{}) {
		if s2.Has(item) {
			inter.Add(item)
		}
	})
	s2.ForEach(func(item interface{}) {
		if s.Has(item) {
			inter.Add(item)
		}
	})

	return inter
}

func (s *set) Union(s2 Set) Set {
	union := New()

	s.readLock()
	defer s.readUnlock()
	s2.readLock()
	defer s2.readUnlock()

	if s2 == nil {
		return s.Copy()
	}

	s.ForEach(func(item interface{}) {
		union.Add(item)
	})
	s2.ForEach(func(item interface{}) {
		union.Add(item)
	})

	return union
}

// Package private methods
func (s *set) rType() reflect.Type {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.rtype
}

func (s *set) readLock() {
	s.mutex.RLock()
}

func (s *set) readUnlock() {
	s.mutex.RUnlock()
}
