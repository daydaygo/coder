package algo

// DP
func fib1(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	return fib1(n-1) + fib1(n-2)
}
func fib2(n int) int {
	if n < 1 {
		return 0
	}
	var memo []int
	return fibHelper(memo, n)
}
func fibHelper(memo []int, n int) int {
	// base case
	if n == 1 || n == 2 {
		return 1
	}
	// 已经计算过
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = fibHelper(memo, n-1) + fibHelper(memo, n-2)
	return memo[n]
}
func fib3(n int) int {
	var dp []int
	// base case
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] // 优化点: 只和前 2 个点有关, 可以优化空间占用
	}
	return dp[n]
}
