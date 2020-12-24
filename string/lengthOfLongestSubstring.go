package string

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
// 示例 1:
//
// 输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
// 输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
// 输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
// Related Topics 哈希表 双指针 字符串 Sliding Window
// 👍 4548 👎 0
//滑动窗口

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0{
		return 0
	}
	m := map[byte]int{}
	res, start := 0, 0
	for end := 0; end < len(s); end ++{
		if _, exit := m[s[end]]; exit{
			start = max(start, m[s[end]] + 1)
		}
		m[s[end]] = end
		res = max(res, end - start + 1)
	}
	return res
}
func max(a int, b int) int {
	if a > b{
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
