package main

import (
	"testing"
)

func TestRemoveElements(t *testing.T) {
	in := []int{1, 2, 6, 3, 4, 5, 6}
	link := createLink(in)
	result := removeElements(link, 6)
	correctResult := createLink([]int{1, 2, 3, 4, 5})

	// dummyHead1, dummyHead2 := &ListNode{}, &ListNode{}
	// dummyHead1.Next = result
	// dummyHead2.Next = correctResult
	cur1 := result
	cur2 := correctResult
	for cur2.Next != nil {
		if cur1.Val != cur2.Val {
			t.Error(cur1.Val, "结果错误")
		} else {
			cur1 = cur1.Next
			cur2 = cur2.Next
		}
	}
}

func createLink(in []int) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead

	for each := range in {
		temp := &ListNode{Val: each}
		// fmt.Println(temp.Next)
		cur.Next = temp
		cur = cur.Next
	}

	return dummyHead.Next
}
