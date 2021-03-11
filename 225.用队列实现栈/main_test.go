package main

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	testStack := Constructor()
	testStack.Push(1)
	testStack.Push(2)
	testStack.Push(3)
	result := testStack.Pop()
	correctResult := 3
	if result != correctResult {
		t.Error(result, "结果错误")
	}
	fmt.Println(testStack)
	result = testStack.Pop()
	correctResult = 2
	if result != correctResult {
		t.Error(result, "结果错误")
	}
	fmt.Println(testStack)
	result = testStack.Pop()
	correctResult = 1
	if result != correctResult {
		t.Error(result, "结果错误")
	}
	fmt.Println(testStack)

	emptyResult := testStack.Empty()
	emptyCorrectResult := true
	if emptyResult != emptyCorrectResult {
		t.Error(emptyResult, "结果错误")
	}
}
