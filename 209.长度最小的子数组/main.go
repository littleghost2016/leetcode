package main

func minSubArrayLen(target int, nums []int) int {
	// left, right := 0, 1
	// minSubArrayLength := len(nums) + 1

	// for right <= len(nums) {
	// 	if sumArray(nums[left:right]) < target {
	// 		right++
	// 	} else {
	// 		if (right - left) < minSubArrayLength {
	// 			minSubArrayLength = right - left
	// 		}
	// 		left++
	// 	}

	// }

	// if minSubArrayLength == len(nums)+1 {
	// 	return 0
	// } else {
	// 	return minSubArrayLength
	// }
	left, right := 0, 0
	sum := 0
	minSubArrayLength := len(nums) + 1
	if len(nums) == 0 {
		return 0
	}

	// 外层循环控制rihgt往右走
	for right < len(nums) {
		sum += nums[right]
		// 内层循环控制left往右走
		for sum >= target {
			minSubArrayLength = min(minSubArrayLength, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}

	if minSubArrayLength == len(nums)+1 {
		return 0
	} else {
		return minSubArrayLength
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
