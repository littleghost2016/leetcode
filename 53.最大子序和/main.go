package main

func maxSubArray(nums []int) int {
	pre := make([]int, len(nums))
	pre[0] = nums[0]

	theMaxSum := pre[0]

	for i := 1; i < len(nums); i++ {
		// 这里是 pre[i-1]+nums[i] 和 nums[i] 进行比较，注意pre[i-1]
		pre[i] = max(pre[i-1]+nums[i], nums[i])

		if pre[i] > theMaxSum {
			theMaxSum = pre[i]
		}

		// fmt.Println(pre)
	}

	// fmt.Println(theMaxSum)
	return theMaxSum
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
