package main

import (
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	result := generateMatrix(3)
	correctResult := [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			if result[i][j] != correctResult[i][j] {
				t.Error(i, j, "结果错误")
			}
		}
	}
}
