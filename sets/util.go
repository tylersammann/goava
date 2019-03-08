package sets

// return s1 - s2
func Difference(s1 Set, s2 Set) Set {
	if s1 == nil {
		return New()
	}

	s1.readLock()
	defer s1.readUnlock()

	diff := s1.Copy()
	if s2 == nil {
		return diff
	}

	s2.readLock()
	defer s2.readUnlock()

	return diff.Remove(s2.Values()...)
}

func Intersection(s1 Set, s2 Set) Set {
	inter := New()

	if s1 == nil || s2 == nil {
		return inter
	}

	s1.readLock()
	defer s1.readUnlock()
	s2.readLock()
	defer s2.readUnlock()

	for _, item := range s1.Values() {
		if s2.Has(item) {
			inter.Add(item)
		}
	}
	for _, item := range s2.Values() {
		if s1.Has(item) {
			inter.Add(item)
		}
	}

	return inter
}

func Union(s1 Set, s2 Set) Set {
	union := New()

	if s1 == nil && s2 == nil {
		return union
	}

	s1.readLock()
	defer s1.readUnlock()
	s2.readLock()
	defer s2.readUnlock()

	if s1 == nil {
		return New(s2.Values()...)
	}
	if s2 == nil {
		return New(s1.Values()...)
	}

	union.Add(s1.Values()...)
	union.Add(s2.Values()...)
	return union
}

// Return true if s1 has the same elements as s2
func Equals(s1 Set, s2 Set) bool {
	if s1 == nil && s2 == nil {
		return true
	}

	if s1 == nil || s2 == nil {
		return false
	}

	s1.readLock()
	defer s1.readUnlock()
	s2.readLock()
	defer s2.readUnlock()

	if len(s1.Values()) != len(s2.Values()) {
		return false
	}

	if s1.RType() != s2.RType() {
		return false
	}

	for _, item := range s1.Values() {
		if !s2.Has(item) {
			return false
		}
	}

	return true
}

// Return true if s1 contains s2
func Contains(s1 Set, s2 Set) bool {
	if s1 == nil && s2 == nil {
		return true
	}

	if s2 == nil {
		return false
	}

	s1.readLock()
	defer s1.readUnlock()
	s2.readLock()
	defer s2.readUnlock()

	// Every set contains the empty set (regardless of type)
	if len(s2.Values()) == 0 {
		return true
	}

	if s1.RType() != s2.RType() {
		return false
	}

	if len(s2.Values()) > len(s1.Values()) {
		return false
	}

	for _, item := range s2.Values() {
		if !s2.Has(item) {
			return false
		}
	}

	return true
}
