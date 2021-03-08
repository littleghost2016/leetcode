package main

type MyQueue struct {
	inStack  []int
	outStack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		this.inToOut()
	}
	x := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		this.inToOut()
	}
	return this.outStack[len(this.outStack)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.inStack) == 0 && len(this.outStack) == 0 {
		return true
	} else {
		return false
	}
}

func (this *MyQueue) inToOut() {
	if len(this.outStack) == 0 {

		// 第一次没写下一行，直接写的for i := 0; i < len(this.inStack); i++
		// 因而会每次计算len(this.inStack)，导致循环次数混乱
		inStackLength := len(this.inStack)
		for i := 0; i < inStackLength; i++ {
			x := this.inStack[len(this.inStack)-1]
			this.inStack = this.inStack[:len(this.inStack)-1]
			this.outStack = append(this.outStack, x)
		}
	}

	// 官方解法
	// for这里会每次重新计算len(q.inStack)
	// for len(q.inStack) > 0 {
	//     q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
	//     q.inStack = q.inStack[:len(q.inStack)-1]
	// }
	// 作者：LeetCode-Solution
	// 链接：https://leetcode-cn.com/problems/implement-queue-using-stacks/solution/yong-zhan-shi-xian-dui-lie-by-leetcode-s-xnb6/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
}
