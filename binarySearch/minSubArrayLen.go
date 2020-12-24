package binarySearch
//给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回
// 0。
//
//
//
// 示例：
//
// 输入：s = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。
//
//
//
//
// 进阶：
//
//
// 如果你已经完成了 O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
//
// Related Topics 数组 双指针 二分查找
// 👍 481 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
import (
	"math"
)
func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	start, end, sum, ans := 0, 0, 0, math.MaxInt32
	for end < len(nums){
		sum += nums[end]
		for sum >= s{
			ans = min(ans, end - start + 1)
			sum -= nums[start]
			start ++
		}
		end ++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
func min(a int, b int) int{
	if a < b{
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)

