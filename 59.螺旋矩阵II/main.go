package main

func generateMatrix(n int) [][]int {

	// 初始化结果切片
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	// 控制循环的圈数（不包括当n为奇数时最里面的一个）
	loop := n >> 1

	// 当n为奇数时，result[middle][middle]为最里面的那一个
	middle := n >> 1

	// 每次递增的数字
	count := 1

	// 每走完一圈，下一圈要走的不属就-2
	// 初始值为1，是因为下面的循环每次从0开始
	offset := 1

	// 每一圈的初始坐标，(0, 0), (1, 1), (2, 2)这样的规律
	startx, starty := 0, 0

	i, j := 0, 0

	for loop > 0 {
		i = startx
		j = starty

		// 每次循环的内容是一样的，但循环的条件不同
		for ; j < starty+n-offset; j++ {
			result[i][j] = count
			count++
		}
		for ; i < startx+n-offset; i++ {
			result[i][j] = count
			count++
		}
		for ; j > starty; j-- {
			result[i][j] = count
			count++
		}
		for ; i > startx; i-- {
			result[i][j] = count
			count++
		}

		startx++
		starty++

		// 左右或者上下都要减去，所以是+2
		offset += 2

		// 圈数每次-1
		loop--
	}

	// 当n为奇数时，单独处理中间的那一个
	if (n & 1) == 1 {
		result[middle][middle] = count
	}

	return result
}
