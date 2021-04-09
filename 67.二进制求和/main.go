package main

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	result := ""
	flag := 0

	// 注意这里的条件都是>=，第一次写的时候写成了>，导致最前面一位没算进去
	for i >= 0 || j >= 0 {
		m, n := 0, 0
		// fmt.Println(i, j)

		// 这里也是>=
		if i >= 0 {
			m = int(a[i] - '0')
			i--
		}
		// 这里也是>=
		if j >= 0 {
			n = int(b[j] - '0')
			j--
		}
		// fmt.Println("m:", m, "n:", n)
		sum := m + n + flag

		switch sum {
		case 0:
			result = "0" + result
			flag = 0
		case 1:
			result = "1" + result
			flag = 0
		case 2:
			result = "0" + result
			flag = 1
		case 3:
			result = "1" + result
			flag = 1
		}
	}

	// 判断是否需要再进一位
	if flag == 1 {
		result = "1" + result
	}

	return result
}
