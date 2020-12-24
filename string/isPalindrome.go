package string

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
// 👍 282 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	//全部转换小写
	s = strings.ToLower(s)
	left, right := 0, len(s) - 1
	for left < right{
		for left < right && !isStringOrNum(s[left]){
			left ++
		}
		for left < right && !isStringOrNum(s[right]){
			right --
		}
		if left < right{
			if s[left] != s[right] {
				return false
			}
			left ++
			right --
		}
	}
	return true
}

func isStringOrNum(ch byte) bool{
	return (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
//leetcode submit region end(Prohibit modification and deletion)

