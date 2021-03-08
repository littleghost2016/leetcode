package main

import "testing"

func TestSumRange02(t *testing.T) {
	myArray := Constructor([]int{-2, 0, 3, -5, 2, -1})

	result := myArray.SumRange(0, 2)
	correctResult := 1
	if result != correctResult {
		t.Error("结果错误")
	}
}

func TestSumRange25(t *testing.T) {
	myArray := Constructor([]int{-2, 0, 3, -5, 2, -1})

	result := myArray.SumRange(2, 5)
	correctResult := -1
	if result != correctResult {
		t.Error("结果错误")
	}
}

func TestSumRange05(t *testing.T) {
	myArray := Constructor([]int{-2, 0, 3, -5, 2, -1})

	result := myArray.SumRange(0, 5)
	correctResult := -3
	if result != correctResult {
		t.Error("结果错误")
	}
}
