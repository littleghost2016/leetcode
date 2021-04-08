package main

// 只有一个数字出现一次，其他均出现了三次，所以要自己写出来一个三进制加法
func singleNumber(nums []int) int {
	result := 0

	// 64是因为go的int是64位的
	for i := 0; i < 64; i++ {
		temp := 0

		// 取出nums中每一个数字的那一位
		for _, each := range nums {
			temp += (each >> i) & 1
		}

		// 或运算，将每一位上的三进制数计入result
		result |= (temp % 3) << i
	}

	return result
}
