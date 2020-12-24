package dp
/**
第300题：最长上升子序列
给定一个无序的整数数组，找到其中最长上升子序列的长度。
示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。

解法：
###LIS可能是连续的，也可能是非连续的。

我们分两种情况进行讨论：

如果nums[i]比前面的所有元素都小，那么dp[i]等于1（即它本身）（该结论正确）
如果nums[i]前面存在比他小的元素nums[j]，那么dp[i]就等于dp[j]+1（该结论错误，比如nums[3]>nums[0]，即9>1,但是dp[3]并不等于dp[0]+1）

我们先初步得出上面的结论，但是我们发现了一些问题。**因为dp[i]前面比他小的元素，不一定只有一个！**
可能除了 nums[j]，还包括 nums[k]，nums[p] 等等等等。所以 dp[i] 除了可能等于 dp[j]+1，还有可能等于 dp[k]+1，dp[p]+1 等等等等。所以我们求 dp[i]，需要找到 dp[j]+1，dp[k]+1，dp[p]+1 等等等等 中的最大值。（我在3个等等等等上都进行了加粗，主要是因为初学者非常容易在这里摔跟斗！这里强调的目的是希望能记住这道题型！） 即：


>
dp[i] = max(dp[j]+1，dp[k]+1，dp[p]+1，.....)
>
只要满足：
>
nums[i] > nums[j]
>
nums[i] > nums[k]
>
nums[i] > nums[p]
>
....

最后，我们只需要找到dp数组中的最大值，就是我们要找的答案。
 */
func lengthOfLIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	res := 1
	for i := 0; i < len(nums); i ++ {
		dp[i] = 1
		for j := 0; j < i; j ++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j] + 1, dp[i])
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}