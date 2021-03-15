package main

import "testing"

func TestRemoveElement(t *testing.T) {
	result := removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	correctResult := 5
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}
