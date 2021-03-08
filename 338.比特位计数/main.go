package main

func countBits1(num int) []int {
	count := make([]int, num+1)
	for i := 0; i <= num; i++ {
		eachCount := 0
		j := i
		for j != 0 {
			j &= (j - 1)
			eachCount++
		}
		count[i] = eachCount
	}
	return count
}

func countBits4(num int) []int {
	count := make([]int, num+1)
	for i := 1; i <= num; i++ {
		count[i] = count[i&(i-1)] + 1
	}
	return count
}

func myEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
