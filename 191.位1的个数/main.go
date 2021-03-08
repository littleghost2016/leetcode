package main

import "fmt"

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

func main() {
	fmt.Println(hammingWeight(11))
}

// n&-n提取数字中二进制最低位1对应的值
// n&n-1消去最低位的1的值。
