package bd

/**
驼峰字符串问题，给定一个驼峰样式的字符串例如“AaABbBcBbcvQv........”->“bc”，
两个一样的字符夹着一个不一样的字符， 处理这个字符串去掉所有的驼峰。
*/

func tuofeng(s string) string {
	if len(s) < 2 {
		return s
	}
	left := 0
	for left+2 < len(s) {
		if s[left] == s[left+2] && s[left] != s[left+1] {
			s = s[left+1:]
			if left >= 2 {
				left -= 2
			} else if left == 1 {
				left -= 1
			}
		} else {
			left += 1
		}
	}
	return s
}
