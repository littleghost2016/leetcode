package main

import "testing"

func TestIsValid(t *testing.T) {
	result := isValid("()[]{}")
	correctResult := true
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}
