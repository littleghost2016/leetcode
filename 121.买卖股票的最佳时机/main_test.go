package main

import "testing"

func TestMaxProfit1(t *testing.T) {
	in := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(in)
	correctResult := 5
	if result != correctResult {
		t.Error(result, "结果错误")
	}

}

func TestMaxProfit2(t *testing.T) {
	in := []int{7, 6, 4, 3, 1}
	result := maxProfit(in)
	correctResult := 0
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}

func TestMaxProfit3(t *testing.T) {
	in := []int{2, 4, 1}
	result := maxProfit(in)
	correctResult := 2
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}
