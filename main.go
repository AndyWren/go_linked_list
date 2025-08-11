package main

import (
	"fmt"
	"reflect"

	"github.com/AndyWren/go-linked-list/linkedlist"
)

func main() {
	var l = linkedlist.NewLinkedList[string]()
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
