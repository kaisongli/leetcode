package arrayandstring

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
// Related Topics 哈希表 双指针 字符串 Sliding Window
// 👍 4322 👎 0

/**
解法:
滑动窗口+双指针

定义start和end双指针，start每次指向不重复字符的开始位置，start每次指向不重复字符的结束位置
如果map中存在当前字符，更新start的位置
此外，无论是否存在字符都需要更新res和end指针的值
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	m := make(map[byte]int)
	start, res := 0, 0
	for end := 0; end < len(s); end++ {
		ch := s[end]
		if _, ok := m[ch]; ok {
			start = max(m[ch]+1, start)
		}
		m[ch] = end
		res = max(res, end-start+1)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
