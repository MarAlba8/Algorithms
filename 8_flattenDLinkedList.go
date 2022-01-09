package main

import "fmt"

var parents_index = map[int]int{
	3: 4,
	8: 2,
}

func FlatList(head *DNode) *DNode {
	walkList(head)
	return head
}

func walkList(head *DNode) *DNode {
	var prev *DNode
	current := head
	for current != nil {
		if current.Child != nil {
			next := current.Next
			tail := walkList(current.Child)
			current.Next = current.Child
			current.Child = nil
			tail.Next = next
		} else {
			prev = current
			current = current.Next
		}
	}
	return prev
}

func CreateDListWithChild(n_sons, start_value int) *DNode {
	/*Creates list of the type
	1---2---3---4---5---6--NULL
	         |
	         7---8---9---10--NULL
	             |
	             11--12--NULL
	*/

	var (
		head *DNode
		prev *DNode
	)

	value := start_value
	for i := 1; i <= n_sons; i++ {
		node := &DNode{
			Value: value,
			Prev:  prev,
			Next:  nil,
			Child: nil,
		}
		if i == 1 {
			head = node
		} else {
			prev.Next = node
		}

		if v, ok := parents_index[value]; ok {
			headChild := CreateDListWithChild(v, start_value+n_sons)
			node.Child = headChild
		}
		value++
		prev = node
	}
	return head
}

func PrintDListValues(head *DNode) {
	node := head
	for node != nil {
		fmt.Print(node.Value)
		fmt.Println("-->")
		if node.Child != nil {
			fmt.Print("\n")
			PrintDListValues(node.Child)
		}
		node = node.Next

	}
}
