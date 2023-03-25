package utilities

import (
	ai "github.com/tsanton/goflake-client/goflake/models/assets/interface"
)

// Stack : a container of objects that are inserted and removed according to the last-in first-out (LIFO) principle
type Stack[T ai.ISnowflakeAsset] struct {
	data []T
}

// IsEmpty : checks whether the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Put : adds an element onto the stack
func (s *Stack[T]) Put(n T) {
	s.data = append(s.data, n)
}

// Get : removes the next element from the queue and returns its value
func (s *Stack[T]) Get() T {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	element := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return element
}
