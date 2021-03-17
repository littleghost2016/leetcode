package main

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func reverseList(head *LinkNode) *LinkNode {
	cur := head
	var pre *LinkNode = nil

	for cur != nil {
		tempNode := cur.Next
		cur.Next = pre

		pre = cur
		cur = tempNode
	}

	return pre
}
