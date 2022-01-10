package jianzhi

/**
剑指 Offer 05. 替换空格
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。



示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."
*/
func replaceSpace(s string) string {
	if len(s) == 0 {
		return ""
	}
	spaceNum := 0
	for _, ch := range s {
		if ch == ' ' {
			spaceNum++
		}
	}
	if spaceNum == 0 {
		return s
	}
	newLen := len(s) + spaceNum*2
	chars := make([]byte, newLen)
	newIdx := newLen - 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			chars[newIdx] = '0'
			newIdx--
			chars[newIdx] = '2'
			newIdx--
			chars[newIdx] = '%'
			newIdx--
		} else {
			if newIdx >= 0 {
				chars[newIdx] = s[i]
				newIdx--
			}
		}
	}
	return string(chars)
}
