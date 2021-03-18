package main

import "testing"

func TestIsHappy19(t *testing.T) {
	result := isHappy(19)
	correctResult := true
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}

func TestIsHappy2(t *testing.T) {
	result := isHappy(2)
	correctResult := false
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}
