package main

func singleNumber(nums []int) int {
	result := 0
	for _, each := range nums {
		result ^= each
	}

	return result
}

// 做过一次了，很简单一遍过
