package string

/**
844. 比较含退格的字符串
给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，请你判断二者是否相等。# 代表退格字符。

如果相等，返回 true ；否则，返回 false 。

注意：如果对空文本输入退格字符，文本继续为空。



示例 1：

输入：s = "ab#c", t = "ad#c"
输出：true
解释：S 和 T 都会变成 “ac”。
示例 2：

输入：s = "ab##", t = "c#d#"
输出：true
解释：s 和 t 都会变成 “”。
*/

//一个字符是否会被删掉，只取决于该字符后面的退格符，而与该字符前面的退格符无关。因此当我们逆序地遍历字符串，就可以立即确定当前字符是否会被删掉。
//
//具体地，我们定义 skip 表示当前待删除的字符的数量。每次我们遍历到一个字符：
//
//若该字符为退格符，则我们需要多删除一个普通字符，我们让 skip 加 1；
//
//若该字符为普通字符：
//
//若 skip 为 0，则说明当前字符不需要删去；
//
//若 skip 不为 0，则说明当前字符需要删去，我们让 skip 减 1。
func backspaceCompare(s string, t string) bool {
	skipS, skipT := 0, 0
	p, q := len(s)-1, len(t)-1
	for p >= 0 || q >= 0 {
		for p >= 0 {
			if s[p] == '#' {
				skipS++
				p--
			} else if skipS > 0 {
				skipS--
				p--
			} else {
				break
			}
		}
		for q >= 0 {
			if t[q] == '#' {
				skipT++
				q--
			} else if skipT > 0 {
				skipT--
				q--
			} else {
				break
			}
		}
		if p >= 0 && q >= 0 {
			if s[p] != t[q] {
				return false
			}
		} else {
			if p >= 0 || q >= 0 {
				return false
			}
		}
		p--
		q--
	}
	return true
}
