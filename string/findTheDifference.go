package string
//给定两个字符串 s 和 t，它们只包含小写字母。
//
// 字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
//
// 请找出在 t 中被添加的字母。
//
//
//
// 示例 1：
//
// 输入：s = "abcd", t = "abcde"
//输出："e"
//解释：'e' 是那个被添加的字母。
//
//
// 示例 2：
//
// 输入：s = "", t = "y"
//输出："y"
//
//
// 示例 3：
//
// 输入：s = "a", t = "aa"
//输出："a"
//
//
// 示例 4：
//
// 输入：s = "ae", t = "aea"
//输出："a"
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 1000
// t.length == s.length + 1
// s 和 t 只包含小写字母
//
// Related Topics 位运算 哈希表
// 👍 169 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findTheDifference(s string, t string) byte {
	m := map[byte]int{}
	for i := 0; i < len(s); i ++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]] = m[s[i]] + 1
		}else {
			m[s[i]] = 1
		}
	}
	for i := 0; i < len(t); i ++ {
		if m[t[i]] == 0 {
			return t[i]
		}else {
			m[t[i]] = m[t[i]] - 1
		}
	}
	return t[len(t) - 1]
}

func findTheDifference2(s string, t string) byte {
	ans := t[len(t) - 1]
	for i := 0; i < len(s); i ++ {
		ans ^= s[i]
		ans ^= t[i]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

