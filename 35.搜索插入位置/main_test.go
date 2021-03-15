package main

import "testing"

func TestSearchInsert1(t *testing.T) {
	result := searchInsert([]int{1, 3, 5, 6}, 0)
	correctResult := 0
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}

func TestSearchInsert2(t *testing.T) {
	result := searchInsert([]int{1, 3, 5, 6}, 2)
	correctResult := 1
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}
