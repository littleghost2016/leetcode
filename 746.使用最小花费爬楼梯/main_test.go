package main

import "testing"

func TestMinCostClimbingStairs1(t *testing.T) {
	result := minCostClimbingStairs([]int{10, 15, 20})
	correctResult := 15
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}

func TestMinCostClimbingStairs2(t *testing.T) {
	result := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	correctResult := 6
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}
