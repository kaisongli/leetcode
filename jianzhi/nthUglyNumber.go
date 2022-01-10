package jianzhi

/**
我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。



示例:

输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
说明:

1 是丑数。
n 不超过1690。
*/
func nthUglyNumber(n int) int {
	idx2, idx3, idx5 := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
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
