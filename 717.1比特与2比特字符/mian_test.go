package main

import "testing"

func TestIsOneBitCharacter(t *testing.T) {
	result := isOneBitCharacter([]int{1, 1, 0})
	correctResult := true
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}
