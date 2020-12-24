package string
/**
第344题：反转字符串
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

示例 1：

输入：["h","e","l","l","o"]
输出：["o","l","l","e","h"]
 */
func reverseString(s []byte) {
	if len(s) < 2 {
		return
	}
	left, right := 0, len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left ++
		right --
	}
}
