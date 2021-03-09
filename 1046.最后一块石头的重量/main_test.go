package main

import "testing"

func TestLastStoneWeight(t *testing.T) {
	result := lastStoneWeight([]int{2, 7, 4, 1, 8, 1})
	correctResult := 1
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}
