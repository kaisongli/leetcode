package bd

/**
264. 丑数 II
编写一个程序，找出第 n 个丑数。

丑数就是质因数只包含 2, 3, 5 的正整数。

示例:

输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
*/
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	idx2, idx3, idx5 := 0, 0, 0
	for i := 1; i < n; i++ {
		n2, n3, n5 := dp[idx2]*2, dp[idx3]*3, dp[idx5]*5
		dp[i] = min(min(n2, n3), n5)
		if dp[i] == n2 {
			idx2++
		}
		if dp[i] == n3 {
			idx3++
		}
		if dp[i] == n5 {
			idx5++
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
