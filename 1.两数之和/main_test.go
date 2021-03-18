package main

import "testing"

func TestTwoSum(t *testing.T) {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	correctResult := []int{0, 1}
	if len(result) != len(correctResult) {
		t.Error(result, "结果错误")
	} else {
		for i := 0; i < len(result); i++ {
			if result[i] != correctResult[i] {
				t.Error(result, "结果错误")
			}
		}
	}
}
