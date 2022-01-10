package string

/**
541. 反转字符串 II
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。

如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。


示例 1：

输入：s = "abcdefg", k = 2
输出："bacdfeg"
示例 2：

输入：s = "abcd", k = 2
输出："bacd"
*/
func reverseStr(s string, k int) string {
	chars := []byte(s)
	n := len(chars)
	for i := 0; i < n; i += 2 * k {
		sub := chars[i:min(i+k, n)]
		reverseString2(sub)
	}
	return string(chars[:])
}

func reverseString2(chars []byte) {
	n := len(chars)
	for i := 0; i < n/2; i++ {
		chars[i], chars[n-i-1] = chars[n-i-1], chars[i]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
