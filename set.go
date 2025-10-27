package main

type Set[T comparable] struct {
	vals map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		vals: make(map[T]struct{}, 0),
	}
}

func (s *Set[T]) Add(item T) {
	s.vals[item] = struct{}{}
}

func (s *Set[T]) ToSlice() []T {
	keys := make([]T, 0, len(s.vals))

	for key := range s.vals {
		keys = append(keys, key)
	}

	return keys
}
