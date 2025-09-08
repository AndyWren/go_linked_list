package stack

import "github.com/AndyWren/go_linked_list/linkedlist"

type Stack[T any] struct {
	list *linkedlist.LinkedList[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{list: linkedlist.NewLinkedList[T]()}
}

func (s *Stack[T]) Push(value T) {
	s.list.AddFirst(value) // Add to front
}

func (s *Stack[T]) Pop() (T, error) {
	return s.list.RemoveFirst() // Remove from front
}

func (s *Stack[T]) Peek() (T, error) {
	return s.list.First() // Look without removing
}

func (s *Stack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}
