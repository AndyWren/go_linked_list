package queue

import "github.com/AndyWren/go-linked-list/linkedlist"

type Queue[T any] struct {
	list *linkedlist.LinkedList[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{list: linkedlist.NewLinkedList[T]()}
}

func (q *Queue[T]) Enqueue(value T) {
	q.list.AddLast(value) // Add to back
}

func (q *Queue[T]) Dequeue() (T, error) {
	return q.list.RemoveFirst() // Remove from front
}

func (q *Queue[T]) Front() (T, error) {
	return q.list.First() // Peek at front
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}
