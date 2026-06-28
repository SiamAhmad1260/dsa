package main

import "fmt"

type Data interface {
	any
}

type Node2 struct {
	data Data
	prev *Node2
	next *Node2
}

type List struct {
	Head *Node2
}

func (l *List) insertBeginning(d Data) {
	newHead := &Node2{data: d}
	if l.Head == nil {
		l.Head = newHead
		l.Head.next = nil
	} else {
		current := l.Head
		l.Head = newHead
		newHead.next = current
		current.prev = newHead
	}
}

func (l *List) insertEnd(d Data) {
	node := &Node2{data: d}
	if l.Head == nil {
		l.Head = node
		return
	} else {
		current := l.Head
		for current.next != nil { //note: for loop stops after false ha been received
			current = current.next
		}
		node.prev = current
		current.next = node
		node.next = nil
	}
}

func (l *List) printList() {
	current := l.Head
	fmt.Println("")
	for current != nil {
		fmt.Printf("%v <-> ", current.data)
		current = current.next
	}
	fmt.Print("nil")
}

func main() {
	l := List{}
	l.insertBeginning(234)
	l.insertBeginning(345)
	l.insertBeginning(98)
	l.printList()
	l.insertEnd("sldj")
	l.insertEnd(598)
	l.printList()
}
