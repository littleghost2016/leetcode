package main

import "testing"

func TestWork(t *testing.T) {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	result := s.GetMin()
	correctResult := -3
	if result != correctResult {
		t.Error(result, "结果错误")
	}
	s.Pop()
	result = s.GetMin()
	correctResult = -2
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}
