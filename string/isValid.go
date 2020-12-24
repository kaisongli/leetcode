package string
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
// 注意空字符串可被认为是有效字符串。
//
// 示例 1:
//
// 输入: "()"
//输出: true
//
//
// 示例 2:
//
// 输入: "()[]{}"
//输出: true
//
//
// 示例 3:
//
// 输入: "(]"
//输出: false
//
//
// 示例 4:
//
// 输入: "([)]"
//输出: false
//
//
// 示例 5:
//
// 输入: "{[]}"
//输出: true
// Related Topics 栈 字符串
// 👍 1991 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	n := len(s)
	if n % 2 == 1{
		return false
	}
	bracketsMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i ++ {
		if bracketsMap[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack) - 1] != bracketsMap[s[i]]{
				return false
			}
			stack = stack[:len(stack) - 1]
		}else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
//leetcode submit region end(Prohibit modification and deletion)
