package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head

	cur := dummyHead

	for cur.Next != nil {
		if cur.Next.Val == val {
			// 以下三行为方案一：
			// 执行用时: 8 ms
			// 内存消耗: 4.6 MB

			temp := cur.Next
			cur.Next = temp.Next
			temp.Next = nil

			// 以下一行为方案二：
			// 执行用时：12 ms, 在所有 Go 提交中击败了20.22%的用户
			// 内存消耗：4.6 MB, 在所有 Go 提交中击败了100.00%的用户
			// 方案二写法简单，但更耗时
			// cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummyHead.Next
}
