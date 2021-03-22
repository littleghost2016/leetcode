package main

func setZeroes(matrix [][]int) {
	// m*n，即n行m列
	n, m := len(matrix), len(matrix[0])

	// 第一列是否有0的标志
	col0 := false

	for i := 0; i < n; i++ {

		// 第一列是否有0
		if matrix[i][0] == 0 {
			col0 = true
		}

		// 每一行的从第二列到最后一列是否有0
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 为了防止第一列的第一个元素被提前更新，需要从最后一行开始，倒序地处理矩阵元素。
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}

		// 判断是否需要处理每一行的第一列
		if col0 == true {
			matrix[i][0] = 0
		}
	}

}
