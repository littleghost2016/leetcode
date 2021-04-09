package main

import "testing"

func TestAddBinary(t *testing.T) {
	result := addBinary("1010", "1011")
	correctResult := "10101"
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}
