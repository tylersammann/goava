package sets

// Difference return s1 - s2
func Difference(s1 Set, s2 Set) Set {
	if s1 == nil {
		return New()
	}

	s1.readLock()
	defer s1.readUnlock()
	s2.readLock()
	defer s2.readUnlock()

	diff := s1.Copy()
	if s2 == nil {
		return diff
	}

	s2.ForEach(func(item interface{}) {
		diff.Remove(item)
	})

	return diff
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

	s1.ForEach(func(item interface{}) {
		if s2.Has(item) {
			inter.Add(item)
		}
	})
	s2.ForEach(func(item interface{}) {
		if s1.Has(item) {
			inter.Add(item)
		}
	})

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
		return s2.Copy()
	}
	if s2 == nil {
		return s1.Copy()
	}

	s1.ForEach(func(item interface{}) {
		union.Add(item)
	})
	s2.ForEach(func(item interface{}) {
		union.Add(item)
	})

	return union
}

// Equals returns true if s1 has the same elements as s2
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

	if s1.Size() != s2.Size() {
		return false
	}

	if s1.rType() != s2.rType() {
		return false
	}

	differentItem := s1.FindFirst(func(item interface{}) bool {
		return !s2.Has(item)
	})

	return differentItem == nil
}

// Contains returns true if s1 contains s2
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
	if s2.Size() == 0 {
		return true
	}

	if s1.rType() != s2.rType() {
		return false
	}

	if s2.Size() > s2.Size() {
		return false
	}

	differentItem := s2.FindFirst(func(item interface{}) bool {
		return !s1.Has(item)
	})

	return differentItem == nil
}
