package main

import (
	"testing"
)

func TestReverseStr(t *testing.T) {
	result := reverseStr("abcdefg", 2)
	correctResult := "bacdfeg"
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}
