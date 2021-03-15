package main

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	result := minSubArrayLen(15, []int{5,1,3,5,10,7,4,9,2,8})
	correctResult := 2
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}
