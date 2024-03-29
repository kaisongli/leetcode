package string

//给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//
//
//
// 示例：
//
// 输入："Let's take LeetCode contest"
//输出："s'teL ekat edoCteeL tsetnoc"
//
//
//
//
// 提示：
//
//
// 在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
//
// Related Topics 字符串
// 👍 262 👎 0
// 字符串转字符数组，按顺序逐个反转

//leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	if len(s) < 2 {
		return s
	}
	//字符串转字符数组
	chars := []byte(s)
	i := 0
	for i < len(chars) {
		start := i
		for i < len(chars) && chars[i] != ' ' {
			i++
		}
		left, right := start, i-1
		for left < right {
			chars[left], chars[right] = chars[right], chars[left]
			left++
			right--
		}
		for i < len(chars) && chars[i] == ' ' {
			i++
		}
	}
	return string(chars)
}
