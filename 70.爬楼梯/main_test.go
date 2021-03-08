package main

import "testing"

func TestClimbStairs(t *testing.T) {
	result := climbStairs(1)
	correctResult := 1
	if result != correctResult {
		t.Error(result, "结果错误")
	}

	result = climbStairs(2)
	correctResult = 2
	if result != correctResult {
		t.Error(result, "结果错误")
	}

	result = climbStairs(3)
	correctResult = 3
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}
