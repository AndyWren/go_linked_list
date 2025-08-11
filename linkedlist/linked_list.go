package linkedlist

import (
	"errors"
	"fmt"
	"iter"
	"reflect"
)

type Node[T any] struct {
	next, prev *Node[T]
	data       T
}

type LinkedList[T any] struct {
	header  *Node[T]
	trailer *Node[T]
	size    int
}

func areEqual(a, b any) (bool, error) {
	if a == nil && b == nil {
		return true, nil
	}
	if a == nil || b == nil {
		return false, nil
	}

	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	if va.Type() != vb.Type() {
		return false, nil
	}

	if !va.Type().Comparable() {
		return false, fmt.Errorf("type %v is not comparable", va.Type())
	}

	return va.Interface() == vb.Interface(), nil
}

func NewLinkedList[T any]() *LinkedList[T] {
	var zero T

	header := &Node[T]{data: zero}
	trailer := &Node[T]{data: zero}

	header.next = trailer
	trailer.prev = header

	return &LinkedList[T]{
		header:  header,
		trailer: trailer,
		size:    0,
	}
}

func (l *LinkedList[T]) retrieveNodeAt(index int) (*Node[T], error) {
	var zero *Node[T]

	if index < 0 || index >= l.size {
		return zero, fmt.Errorf("index %d out of range [0, %d]", index, l.size)
	}

	if index == 0 {
		return l.header.next, nil
	}

	current := l.header.next
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current, nil
}

func (l *LinkedList[T]) Get(index int) (T, error) {
	var zero T

	n, err := l.retrieveNodeAt(index)

	if err != nil {
		return zero, err
	}

	return n.data, nil
}

func (l *LinkedList[T]) SetAtIndex(index int, value T) error {
	node, err := l.retrieveNodeAt(index)
	if err != nil {
		return err
	}
	node.data = value
	return nil
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) AddFirst(value T) {
	firstNode := l.header.next
	newNode := &Node[T]{
		data: value,
		prev: l.header,
		next: firstNode,
	}
	firstNode.prev = newNode
	l.header.next = newNode
	l.size++
}

func (l *LinkedList[T]) AddLast(value T) {
	lastNode := l.trailer.prev
	newNode := &Node[T]{
		data: value,
		prev: lastNode,
		next: l.trailer,
	}

	lastNode.next = newNode
	l.trailer.prev = newNode
	l.size++
}

func (l *LinkedList[T]) AddAtIndex(index int, value T) error {
	nodeAtIndex, err := l.retrieveNodeAt(index)
	if err != nil {
		return err
	}

	newNode := &Node[T]{
		data: value,
		prev: nodeAtIndex.prev,
		next: nodeAtIndex,
	}

	nodeAtIndex.prev.next = newNode
	nodeAtIndex.prev = newNode
	l.size++
	return nil
}

func (l *LinkedList[T]) RemoveFirst() (T, error) {
	var zero T

	if l.header.next == l.trailer {
		return zero, errors.New("list is empty")
	}

	firstNode := l.header.next

	l.header.next = firstNode.next
	firstNode.next.prev = l.header

	l.size--
	return firstNode.data, nil
}

func (l *LinkedList[T]) RemoveLast() (T, error) {
	var zero T
	if l.trailer.prev == l.header {
		return zero, errors.New("list is empty")
	}

	lastNode := l.trailer.prev
	l.trailer.prev = lastNode.prev
	lastNode.prev.next = l.trailer
	l.size--
	return lastNode.data, nil
}

func (l *LinkedList[T]) RemoveAtIndex(index int) (T, error) {
	var zero T
	nodeAtIndex, err := l.retrieveNodeAt(index)
	if err != nil {
		return zero, err
	}

	nodeAtIndex.prev.next = nodeAtIndex.next
	nodeAtIndex.next.prev = nodeAtIndex.prev
	l.size--
	return nodeAtIndex.data, nil
}

func (l *LinkedList[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		current := l.header.next
		for current != l.trailer {
			if !yield(current.data) {
				return
			}
			current = current.next
		}
	}
}

func (l *LinkedList[T]) Backward() iter.Seq[T] {
	return func(yield func(T) bool) {
		current := l.trailer.prev
		for current != l.header {
			if !yield(current.data) {
				return
			}
			current = current.prev
		}
	}
}

func (l *LinkedList[any]) IndexOf(value any) (int, error) {
	current := l.header.next
	for i := 0; i < l.size; i++ {
		equal, err := areEqual(current.data, value)
		if err != nil {
			return -1, err
		}
		if equal {
			return i, nil
		}
		current = current.next
	}
	return -1, nil
}

func (l *LinkedList[any]) Contains(value any) (bool, error) {
	index, err := l.IndexOf(value)
	if err != nil {
		return false, err
	}
	return index != -1, nil
}

func (l *LinkedList[any]) RemoveValue(value any) (bool, error) {
	current := l.header.next
	for current != l.trailer { // Better loop condition
		equal, err := areEqual(current.data, value)
		if err != nil {
			return false, err
		}
		if equal {
			current.prev.next = current.next
			current.next.prev = current.prev
			l.size--
			return true, nil
		}
		current = current.next
	}
	return false, nil
}

func (l *LinkedList[T]) Clear() {
	l.header.next = l.trailer
	l.trailer.prev = l.header
	l.size = 0
}

func (l *LinkedList[T]) First() (T, error) {
	var zero T
	if l.header.next == l.trailer {
		return zero, errors.New("list is empty")
	}

	return l.header.next.data, nil
}

func (l *LinkedList[T]) Last() (T, error) {
	var zero T
	if l.trailer.prev == l.header {
		return zero, errors.New("list is empty")
	}

	return l.trailer.prev.data, nil
}
