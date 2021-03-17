package main

type LinkNode struct {
	Val  int
	Next *LinkNode
}

type MyLinkedList struct {
	Head   *LinkNode
	Tail   *LinkNode
	Length int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	// 初始化MyLinkedList记得给Head和Tail也初始化
	// 否则会报错invalid memory address or nil pointer dereference
	return MyLinkedList{
		Head:   &LinkNode{},
		Tail:   &LinkNode{},
		Length: 0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	// 先判断索引是否有效
	if 0 <= index && index < this.Length {
		// 注意这里的cur为Head.Next，后面的add***为Head
		// 是因为Add的时候要从要A增加位置的前一个开始操作，而Get则在目标位置即可
		cur := this.Head.Next
		for i := 0; i < index; i++ {
			cur = cur.Next
		}
		return cur.Val
	} else {
		return -1
	}
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	tempNode := &LinkNode{
		Val: val,
	}

	// 当头节点后面没有节点时
	if this.Head.Next == nil {
		this.Head.Next = tempNode
		// 注意这时候要指定尾节点
		this.Tail = tempNode
	} else {
		tempNode.Next = this.Head.Next
		this.Head.Next = tempNode
	}

	this.Length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	// 相当于AddAtHead
	if this.Length == 0 {
		this.AddAtHead(val)
	} else {
		tempNode := &LinkNode{
			Val: val,
		}
		this.Tail.Next = tempNode
		// 重新指定尾节点
		this.Tail = tempNode
		this.Length++
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
	} else if index == this.Length {
		this.AddAtTail(val)
	} else if index > this.Length {
		return
	} else {
		cur := this.Head
		for i := 0; i < index; i++ {
			cur = cur.Next
		}
		tempNode := &LinkNode{
			Val: val,
		}
		tempNode.Next = cur.Next
		cur.Next = tempNode
		// 这里的不能放到if-else的判断之外，刚开始放到外面了，导致自增了2次
		this.Length++
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Length || this.Length == 0 {
		return
	} else {
		cur := this.Head
		for i := 0; i < index; i++ {
			cur = cur.Next
		}
		if cur.Next.Next != nil {
			tempNode := cur.Next
			cur.Next = tempNode.Next
			tempNode.Next = nil
		} else {
			cur.Next = nil
			// 重新指定尾节点
			this.Tail = cur
		}
		this.Length--
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
