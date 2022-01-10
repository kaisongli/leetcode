package dp

//你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的
//房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,3,2]
//输出：3
//解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3,1]
//输出：4
//解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
//     偷窃到的最高金额 = 1 + 3 = 4 。
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 1000
//
// Related Topics 动态规划
// 👍 691 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	//抢第一家或抢最后一家
	return max(robRange(nums, 0, len(nums)-2), robRange(nums, 1, len(nums)-1))
}

func robRange(nums []int, start, end int) int {
	dp := make([]int, len(nums)+2)
	ans := 0
	for i := end; i >= start; i-- {
		//抢下一家，抢当前这家+下下家
		dp[i] = max(dp[i+1], nums[i]+dp[i+2])
		ans = max(ans, dp[i])
	}
	return ans
}

func rob_2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	return max(robRange(nums, 0, n-2), robRange(nums, 1, n-1))
}

func robRange_2(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}
	dp := make([]int, len(nums))
	dp[start] = nums[start]
	dp[start+1] = max(nums[start], nums[start+1])
	for i := start + 2; i <= end; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[end]
}

//leetcode submit region end(Prohibit modification and deletion)
