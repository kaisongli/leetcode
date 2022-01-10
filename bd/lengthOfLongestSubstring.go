package bd

/**
剑指 Offer 48. 最长不含重复字符的子字符串
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	m := make(map[byte]int)
	start, ans := 0, 0
	for end := 0; end < len(s); end++ {
		if _, ok := m[s[end]]; ok {
			start = max(start, m[s[end]]+1)
		}
		m[s[end]] = end
		ans = max(ans, end-start+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
