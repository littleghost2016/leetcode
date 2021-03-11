package main

import "testing"

func TestReverse(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	reverse(a)
	correctResult := []int{5, 4, 3, 2, 1}
	if !equal(a, correctResult) {
		t.Error(a, "结果错误")
	}
}

func TestRotate(t *testing.T) {
	a := []int{-1, -100, 3, 99}
	k := 2
	rotate(a, k)
	correctResult := []int{3, 99, -1, -100}
	if !equal(a, correctResult) {
		t.Error(a, "结果错误")
	}
}

func equal(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
