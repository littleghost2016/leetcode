package main

// 用时0ms
func isHappy(n int) bool {
	hashMap := make(map[int]bool, 0)

	for n != 1 {
		sum := 0
		tempN := n
		for tempN != 0 {
			a := tempN % 10
			sum += a * a
			tempN = (tempN - a) / 10
		}

		if sum == 1 {
			return true
		} else if _, ok := hashMap[sum]; ok {
			return false
		} else {
			hashMap[sum] = true
		}

		n = sum
	}

	return true
}

// 用时4ms
// func isHappy1(n int) bool {
// 	// 检验是否存在
// 	hashMap := make(map[int]bool, 0)

// 	// 满足快乐数条件则退出
// 	for n != 1 {
// 		sum := 0
// 		tempN := n

// 		for tempN != 0 {
// 			a := tempN % 10
// 			sum += a * a
// 			tempN = (tempN - a) / 10
// 		}

// 		n = sum

// 		if _, ok := hashMap[sum]; ok {
// 			return false
// 		} else {
// 			hashMap[sum] = true
// 		}
// 	}

// 	return true
// }
