package string
//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
// 示例 1：
//
// 输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//
//
// 示例 2：
//
// 输入: "cbbd"
//输出: "bb"
//
// Related Topics 字符串 动态规划
// 👍 2871 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	if len(s) < 2{
		return s
	}
	var res string
	for i := 0; i < len(s); i ++{
		for j := i + 1; j <= len(s); j ++{
			if isPalindrome(s[i: j]) && j - i + 1 > len(res){
				res = s[i: j]
			}
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

