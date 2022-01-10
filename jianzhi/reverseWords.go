package jianzhi

/**
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。



示例 1：

输入: "the sky is blue"
输出: "blue is sky the"
示例 2：

输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
*/

import "strings"

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	words := strings.Split(s, " ")
	var ans []string
	for i := len(words) - 1; i >= 0; i-- {
		str := strings.TrimSpace(words[i])
		if len(str) > 0 {
			ans = append(ans, str)
		}
	}
	return strings.Join(ans, " ")
}
