package main

func reverse(a []int) {
	length := len(a)
	for i := 0; i < (length / 2); i++ {
		a[i], a[length-1-i] = a[length-1-i], a[i]
	}
}

func rotate(nums []int, k int) {
	// 简化运算
	// 当nums=[-1], k=2时会出现slice越界
	k = k % len(nums)

	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
