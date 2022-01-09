package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type DNode struct {
	Value int
	Prev  *DNode
	Next  *DNode
	Child *DNode
}

func CreateLinkedList(n int) *Node {
	i := 1
	prevNode := &Node{}
	head := &Node{}
	for i <= n {
		node := Node{
			Value: i,
			Next:  nil,
		}
		if i == 1 {
			head = &node
		}
		prevNode.Next = &node
		prevNode = &node
		i++
	}
	return head
}

func PrintListValues(head *Node) {
	node := head
	for node != nil {
		fmt.Println(node.Value)
		node = node.Next
	}
}

/*
func ReverseLinkedList(head *Node, newList *Node) *Node {
	if head == nil {
		return newList
	}

	var restOfList *Node
	currentNode := head.Next
	if currentNode != nil {
		restOfList = currentNode.Next
		currentNode.Next = head
		head.Next = newList
		newList = ReverseLinkedList(restOfList, currentNode)
	} else {
		head.Next = newList
		newList = head
	}
	return newList
}
*/
