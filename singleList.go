package dsa

import (
	"fmt"
)

type Node struct {
	data Data
	next *Node
}

type List_ struct {
	Head *Node
}

type ListProcess interface {
	insertBeginning(d Data)
	insertEnd(d Data)
	PrintList()
	index(i uint) Data
	len() uint32
	search(d Data) []int
	deleteNode(i uint)
}

// *var -> data inside var
// *type -> pointer of type
// & -> address of var

// Adds a new data node in the end
func (ll *List_) insertEnd(d Data) {
	nextNode := &Node{data: d, next: nil}
	if ll.Head == nil {
		ll.Head = nextNode
	} else {
		current := ll.Head
		for current.next != nil {
			current = current.next
		}
		current.next = nextNode
	}
}

// Adds a new data node at the beginning as the new head
func (ll *List_) insertBeginning(d Data) {
	newHead := &Node{data: d, next: nil}
	if ll.Head == nil {
		ll.Head = newHead
	} else {
		current := ll.Head
		ll.Head = newHead
		newHead.next = current
	}
}

func (ll *List) insertAfter(d Data, n uint) {
	if n == 1 {
		ll.insertBeginning(d)
		return
	}
	i := 0
	node := &Node2{data: d}
	current := ll.Head
	for current != nil && i != n+1 {
		i++
		if current == nil || current.next == nil {
			fmt.Println("Out of bound")
			return
		}
		current = current.next
	}
	nextnode := current.next.next
	current.next = node
	node.next = nextnode
}

// Prints out the entire list
func (ll *List_) PrintList() {
	node := ll.Head
	for node != nil {
		fmt.Printf("%v -> ", node.data)
		node = node.next
	}
	fmt.Printf("nil\n")

}

// returns the  data in the given index
func (ll *List_) index(i uint) Data {
	current := ll.Head
	var j uint = 1
	for j != i {
		current = current.next
		j = j + 1
	}
	return current.data
}

// gives an integer about how many nodes a list contains
func (ll *List_) len() uint32 {
	var i uint32 = 0
	current := ll.Head
	for current != nil {
		i = i + 1
		current = current.next
	}
	return i
}

// returns a slice of indexes in which the data can be found
func (ll *List_) search(d Data) []int {
	var result []int
	current := ll.Head
	index := 1
	for current.next != nil {
		if d == current.data {
			result = append(result, index)
		}
		index = index + 1
		current = current.next
	}
	return result
}

func (ll *List_) deleteNode(i uint) {
	current := ll.Head
	if i < 1 || ll.Head == nil {
		return
	}
	if i == 1 {
		current = current.next
	}
	j := 1
	for current != nil && j != i {
		current = current.next
		j++
	}

	if current == nil || current.next == nil {
		return
	}
	current.next = current.next.next
}

/*
func main() {
	link := List_{}
	link.insertBeginning(12)
	link.insertBeginning(907)
	link.insertBeginning(23.567)
	link.insertBeginning("aldskfoaeioi")
	link.insertBeginning(12)
	link.insertBeginning(12)
	link.insertBeginning(12)

	link.PrintList()
	link.insertEnd(2333523)
	link.PrintList()

	fmt.Println("search result:", link.search(907))
	fmt.Println("search result:", link.search("aldskfoaeioi"))
	fmt.Println("search result:", link.search(67))
	fmt.Println("search result:", link.search(12))

	new := List_{}
	for i := 1; i <= 50; i = i + 5 {
		new.insertBeginning(i)
	}
	new.PrintList()
	for j := 100; j > 50; j = j - 5 {
		new.insertEnd(j)
	}

	new.PrintList()
	fmt.Println(new.len())
	new.deleteNode(7)
	new.PrintList()
	fmt.Println(new.len())
}
*/
