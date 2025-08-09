package main

import (
	"errors"
	"fmt"
	"iter"
	"reflect"
)

func mod(x int, y int) int {
	for x >= y {
		x -= y
	}
	return x
}

type Node[T comparable] struct {
	next, prev *Node[T]
	data       T
}

type LinkedList[T comparable] struct {
	header  *Node[T]
	trailer *Node[T]
	size    int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
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

func (l *LinkedList[T]) Contains(value T) bool {
	return l.IndexOf(value) != -1
}

func (l *LinkedList[T]) IndexOf(value T) int {
	if l.size == 0 {
		return -1
	}

	current := l.header.next
	for i := 0; i < l.size; i++ {
		if current.data == value {
			return i
		}
		current = current.next
	}
	return -1
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

func main() {
	fmt.Println(mod(32, 12))
	var l = NewLinkedList[string]()
	l.AddLast("world")
	fmt.Println(l.Get(0))
	l.AddFirst("hello")
	fmt.Println(l.Get(0))
	for v := range l.All() {
		fmt.Println(v)
	}
	err := l.SetAtIndex(3, "broke")
	if err != nil {
		fmt.Println(err)
	}
	t := "frank"
	vt := reflect.ValueOf(t)
	fmt.Println(vt.Type())
	fmt.Println(vt.Type().Comparable())
	fmt.Println(vt.Interface())

}
