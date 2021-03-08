package main

import "testing"

func TestStack(t *testing.T) {
	myQueue := Constructor()
	myQueue.Push(1)
	t.Log(myQueue)
	myQueue.Push(2)
	t.Log(myQueue)
	myQueue.Peek()
	t.Log(myQueue)
	myQueue.Pop()
	t.Log(myQueue)
	myQueue.Empty()
	t.Log(myQueue)
}
