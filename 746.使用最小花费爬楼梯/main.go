package main

func minCostClimbingStairs(cost []int) int {
	// 注意这个next
	pre, cur, next := 0, 0, 0

	for i := 2; i <= len(cost); i++ {
		next = myMin(pre+cost[i-2], cur+cost[i-1])
		pre = cur
		cur = next
	}

	return next
}

func myMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
