package main

import "testing"

func TestSingleNumber(t *testing.T) {
	result := singleNumber([]int{2, 2, 2, 3})
	correctResult := 3
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}
