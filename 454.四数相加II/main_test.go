package main

import "testing"

func TestFourSumCount(t *testing.T) {
	result := fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2})
	correctResult := 2
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}
