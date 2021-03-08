package main

import "testing"

func TestMaxSubArray(t *testing.T) {
	in := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubArray(in)
	correctResult := 6
	if result != correctResult {
		t.Log(result)
		t.Error("结果错误")
	}

	in = []int{1}
	result = maxSubArray(in)
	correctResult = 1
	if result != correctResult {
		t.Log(result)
		t.Error("结果错误")
	}

	in = []int{-1, -2}
	result = maxSubArray(in)
	correctResult = -1
	if result != correctResult {
		t.Log(result)
		t.Error("结果错误")
	}
}
