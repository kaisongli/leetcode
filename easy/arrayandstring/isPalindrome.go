package arrayandstring

import "strings"

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
// 说明：本题中，我们将空字符串定义为有效的回文串。
//
// 示例 1:
//
// 输入: "A man, a plan, a canal: Panama"
//输出: true
//
//
// 示例 2:
//
// 输入: "race a car"
//输出: false
//
// Related Topics 双指针 字符串
// 👍 275 👎 0
/**
1、清除左右空格，全部切换为小写字母
2、双指针跳过非字母和数字字符
3、如果左指针和右指针的值不想等，直接返回false
*/

func isPalindrome(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isCharOrNum(s[left]) {
			left++
		}
		for left < right && !isCharOrNum(s[right]) {
			right--
		}
		if left < right && s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isCharOrNum(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}
