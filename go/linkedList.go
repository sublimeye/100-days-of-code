package main

import (
	"fmt"
)

type Node struct {
	next  *Node
	list  *List
	value interface{}
}

type List struct {
	head   Node
	length int
}

func NewList() *List {
	return new(List).Init()
}

func (l *List) Init() *List {
	l.head.next = &l.head
	l.length = 0
	return l
}

func (l *List) Insert(element, at *Node) *Node {
	next := at.next
	at.next = element
	element.next = next
	element.list = l
	l.length++
	return element
}

func (l *List) InsertValue(value interface{}, at *Node) *Node {
	return l.Insert(&Node{value: value}, at)
}

func (l *List) PushFront(value interface{}) *Node {
	return l.InsertValue(value, &l.head)
}

func (l *List) Length() int {
	return l.length
}

// Init
// next
// prev
// addLast
// length
// delete

// addFirst
// insertAfter
// insertBefore

func main() {
	my := NewList()
	fmt.Println(my.Length())
	my.PushFront("test")
	my.PushFront("one")
	my.PushFront("two")
	my.PushFront("three")
	fmt.Println(my)
	fmt.Println(my.Length())
}
