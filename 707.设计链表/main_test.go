package main

import (
	"fmt"
	"testing"
)

func TestProcess1(t *testing.T) {
	myLinkedList := Constructor()
	myLinkedList.AddAtHead(1)
	myLinkedList.AddAtTail(3)
	myLinkedList.AddAtIndex(1, 2)
	// fmt.Printf("%#v\n", myLinkedList.Head)

	// correctMyLinkedList := createMyLinkedList([]int{1, 2, 3})
	// result := compare(&myLinkedList, correctMyLinkedList)
	// if result != true {
	// 	t.Error(&myLinkedList, "结果错误")
	// }
	myLinkedList.DeleteAtIndex(1)
	printLinkNodeHead(myLinkedList.Head)
}

func TestProcess2(t *testing.T) {
	myLinkedList := Constructor()
	myLinkedList.AddAtHead(0)
	myLinkedList.AddAtIndex(1, 4)
	myLinkedList.AddAtTail(8)
	myLinkedList.AddAtHead(5)
	myLinkedList.AddAtIndex(4, 3)
	myLinkedList.AddAtTail(0)
	myLinkedList.AddAtTail(5)
	myLinkedList.AddAtIndex(6, 3)
	myLinkedList.DeleteAtIndex(7)
	myLinkedList.DeleteAtIndex(5)
	myLinkedList.AddAtTail(4)
	// printLinkNodeHead(myLinkedList.Head)
	// fmt.Println(myLinkedList.Tail.Val, myLinkedList.Length)
}

func TestProcess3(t *testing.T) {
	myLinkedList := Constructor()
	myLinkedList.AddAtHead(4)
	myLinkedList.Get(1)
	printLinkNodeHead(myLinkedList.Head)
	fmt.Println(myLinkedList.Tail.Val, myLinkedList.Length)
	myLinkedList.AddAtHead(1)
	myLinkedList.AddAtHead(5)
	myLinkedList.DeleteAtIndex(3)
	myLinkedList.AddAtHead(7)
	myLinkedList.Get(3)
	myLinkedList.Get(3)
	myLinkedList.Get(3)
	myLinkedList.AddAtHead(1)
	myLinkedList.DeleteAtIndex(4)
}

func compare(a, b *MyLinkedList) bool {
	if a.Length != b.Length {
		return false
	}

	cur1, cur2 := a.Head, b.Head
	for cur1.Next != nil && cur2.Next != nil {
		// fmt.Println(cur1.Val, cur2.Val)
		if cur1.Val != cur2.Val {
			// t.Error(cur1.Val, cur2.Val, "结果错误")
			return false
		} else {
			cur1 = cur1.Next
			cur2 = cur2.Next
		}
	}

	return true
}

func createLinkedList(in []int) *LinkNode {
	dummyHead := &LinkNode{}
	cur := dummyHead
	for each := range in {
		newNode := &LinkNode{
			Val: each,
		}
		cur.Next = newNode
		cur = cur.Next
	}

	return dummyHead.Next
}

func createMyLinkedList(in []int) *MyLinkedList {
	newLinkedList := &MyLinkedList{
		Head:   &LinkNode{},
		Tail:   &LinkNode{},
		Length: len(in),
	}

	cur := newLinkedList.Head

	for i := 0; i < len(in); i++ {
		newLinkNode := &LinkNode{Val: in[i]}
		cur.Next = newLinkNode
		if i == (len(in) - 1) {
			newLinkedList.Tail = newLinkNode
		}
	}

	return newLinkedList
}

func printLinkNodeHead(in *LinkNode) {
	cur := in.Next
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}
