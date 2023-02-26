package main

import "fmt"

type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}
type DoublyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (ll *DoublyLinkedList) RemoveAtFirst() {
	if ll.Head == nil {
		fmt.Println("Empty")
	}
	ll.Head = ll.Head.Next
	ll.Tail = ll.Head
}

func (ll *DoublyLinkedList) AddAtFirst(data int) {
	newNode := &Node{Value: data}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Head.Previous = newNode
		newNode.Next = ll.Head
		ll.Head = newNode
	}
}

func (ll *DoublyLinkedList) Display() {
	current := ll.Head
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println("")
}
func main() {
	List := DoublyLinkedList{}
	//Adding at first
	List.AddAtFirst(1)
	List.AddAtFirst(2)
	List.AddAtFirst(3)
	List.AddAtFirst(4)

	//Removing at first
	List.RemoveAtFirst()

	List.Display()
}
