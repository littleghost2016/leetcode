package main

import (
	"testing"
)

var correctResult = []int{0, 1, 1, 2, 1, 2}

func TestCountBits1(t *testing.T) {
	result := countBits1(5)
	if !myEqual(result, correctResult) {
		t.Error("结果错误")
	}
}

func TestCountBits4(t *testing.T) {
	result := countBits4(5)
	if !myEqual(result, correctResult) {
		t.Error("结果错误")
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
