package main

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
