package main

import "testing"

func TestCalculate(t *testing.T) {
	result := calculate("11+1+1")
	correctResult := 13
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}
