package main

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	// 此时区间为左闭右开，即[left, right)

	middle := 0

	for left < right {
		middle = left + (right-left)>>1
		// 小于，left要middle+1
		if nums[middle] < target {
			left = middle + 1

			// 大于，right直接为middle
		} else if nums[middle] > target {
			right = middle
		} else {
			return middle
		}
	}

	// 返回右
	return right
}
