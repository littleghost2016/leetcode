package main

func removeElement(nums []int, val int) int {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		j++
	}

	// 以下三行用作减少内存消耗
	if len(nums) != 0 {
		nums = nums[:i]
	}

	return i
}
