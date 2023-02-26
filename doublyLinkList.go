package main

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
