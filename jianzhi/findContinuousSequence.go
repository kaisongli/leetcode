package jianzhi

/**
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。



示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]
*/
func findContinuousSequence(target int) [][]int {
	left, right, sum := 1, 1, 0
	var ans [][]int
	for left <= target/2+1 {
		if sum < target {
			sum += right
			right++
		} else if sum > target {
			sum -= left
			left++
		} else {
			res := make([]int, right-left)
			for i := left; i < right; i++ {
				res[i-left] = i
			}
			ans = append(ans, res)
			sum -= left
			left++
		}
	}
	return ans
}
