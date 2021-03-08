package main

import "testing"

func TestIsSubsequence1(t *testing.T) {
	result := isSubsequence("abc", "ahbgdc")
	correctResult := true
	if result != correctResult {
		t.Error("结果错误")
	}
}

func TestIsSubsequence2(t *testing.T) {
	result := isSubsequence("", "ahbgdc")
	correctResult := true
	if result != correctResult {
		t.Error("结果错误")
	}
}
