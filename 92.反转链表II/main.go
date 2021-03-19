package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	pre := dummyHead

	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next

		// 注意指向的顺序
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return dummyHead.Next
}
