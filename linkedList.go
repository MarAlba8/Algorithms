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

func ReversePortion(head *Node, m, n int) *Node {
	if head == nil {
		return nil
	}

	if head.Next == nil || m == n {
		return head
	}

	var newHead *Node
	var prev *Node
	count := 1
	current := head
	for count < m {
		prev = current
		current = current.Next
		count++
	}
	newHead = ReverseLinkedList(current, n-count+1)
	if prev != nil {
		prev.Next = newHead
		return head
	} else {
		return newHead
	}
}

func ReverseLinkedList(head *Node, n int) *Node {
	count := 1
	if count > n {
		return head
	}
	var (
		next *Node
		prev *Node
	)
	current := head

	for count <= n {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
		count++
	}
	head.Next = current
	return prev
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
	PrintListValues(head)
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
