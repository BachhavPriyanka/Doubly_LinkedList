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

func (ll *DoublyLinkedList) RemoveAtLast() {
	if ll.Head == nil {
		fmt.Println("List is Empty")
		return
	} else if ll.Head == ll.Tail {
		ll.Head = nil
		ll.Tail = nil
	} else {
		ll.Tail = ll.Tail.Previous
		ll.Tail.Next = nil
	}
	ll.Size--
}
func (ll *DoublyLinkedList) AddAtLast(data int) {
	newNode := &Node{Value: data}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		current := ll.Head

		for current.Next != nil {
			current = current.Next
		}
		newNode.Previous = current
		current.Next = newNode
		ll.Tail = newNode
	}
	ll.Size++

}

func (ll *DoublyLinkedList) RemoveAtFirst() {
	if ll.Head == nil {
		fmt.Println("Empty")
	}
	ll.Head = ll.Head.Next
	ll.Tail = ll.Head

	ll.Size--
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
	ll.Size++
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

	//Adding node at last
	List.AddAtLast(23)

	//Removing node from last
	List.RemoveAtLast()
	List.RemoveAtLast()
	List.Display()
}
