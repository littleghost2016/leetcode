package main

func fourSumCount(A []int, B []int, C []int, D []int) int {
	valueMap := make(map[int]int, 0)
	for _, eachA := range A {
		for _, eachB := range B {
			valueMap[eachA+eachB]++
		}
	}

	count := 0

	for _, eachC := range C {
		for _, eachD := range D {
			if value, ok := valueMap[0-(eachC+eachD)]; ok {
				count += value
			}
		}
	}

	return count
}
