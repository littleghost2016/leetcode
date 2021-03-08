package main

func climbStairs(n int) int {
	p, q, r := 0, 1, 0
	for i := 0; i < n; i++ {
		r = p + q
		p = q
		q = r
	}

	return r
}

// 状态转移方程：f(x)=f(x−1)+f(x−2)
