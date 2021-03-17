package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	// 快慢指针
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// 指向快指针
			index1 := fast
			// 指向头节点
			index2 := head
			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next
			}
			return index2
		}
	}

	return nil
}
