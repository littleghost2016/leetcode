package main

import "math"

type MinStack struct {
	stack    []int
	minStack []int
	length   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
		length:   0,
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	// 注意下面这一行的this.minStack[this.length]
	// 因为初始化minStack的时候已经压入了一个math.MaxInt64，所以栈顶的索引是length而不是length-1
	this.minStack = append(this.minStack, min(val, this.minStack[this.length]))

	// 别忘了长度增加
	this.length++
}

func (this *MinStack) Pop() {
	// 判断栈是否为空
	if this.length > 0 {
		this.stack = this.stack[:this.length-1]

		// 下面这一行的[:this.length]同Push里面的解释
		this.minStack = this.minStack[:this.length]

		// 别忘了长度减少
		this.length--
	}
}

func (this *MinStack) Top() int {
	return this.stack[this.length-1]
}

func (this *MinStack) GetMin() int {

	// 下面这一行的this.minStack[this.length]同Push里面的解释
	return this.minStack[this.length]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
