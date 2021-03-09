package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	result := removeDuplicates("abbaac")
	correctResult := "ac"
	if result != correctResult {
		t.Error(len(result), result, "结果错误")
	}
}
