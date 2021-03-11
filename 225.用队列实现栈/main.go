package main

type MyStack struct {
	queue1     []int
	queue2     []int
	queue2Flag bool
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if !this.queue2Flag {
		this.queue1 = append(this.queue1, x)
	} else {
		this.queue2 = append(this.queue2, x)
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	} else {
		if !this.queue2Flag {
			temp := this.queue1[len(this.queue1)-1]
			// copy(this.queue2, this.queue1[:len(this.queue1)-1])
			// 此处可以直接交换，不用copy
			// copy的时候以两个切片中长度小的那个为准
			// 即copy([]int{1,2,3}, []int{4,5})的结果为[]int{4,5,3}
			this.queue1, this.queue2 = this.queue2, this.queue1
			// fmt.Println(this)
			this.queue2 = this.queue2[:len(this.queue2)-1]
			this.queue2Flag = true
			return temp
		} else {
			temp := this.queue2[len(this.queue2)-1]
			// copy(this.queue2, this.queue1[:len(this.queue1)-1])
			this.queue1, this.queue2 = this.queue2, this.queue1
			// fmt.Println(this)
			this.queue1 = this.queue1[:len(this.queue1)-1]
			this.queue2Flag = false
			return temp
		}
	}
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if !this.queue2Flag {
		return this.queue1[len(this.queue1)-1]
	} else {
		return this.queue2[len(this.queue2)-1]
	}
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	// fmt.Println(this.queue1, this.queue2)
	if !this.queue2Flag {
		if len(this.queue1) != 0 {
			return false
		} else {
			return true
		}
	} else {
		if len(this.queue2) != 0 {
			return false
		} else {
			return true
		}
	}
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
