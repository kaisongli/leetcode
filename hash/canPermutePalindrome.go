package hash

//给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。
//
// 回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。
//
// 回文串不一定是字典当中的单词。
//
//
//
// 示例1：
//
// 输入："tactcoa"
//输出：true（排列有"tacocat"、"atcocta"，等等）
//
//
//
// Related Topics 哈希表 字符串
// 👍 40 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func canPermutePalindrome(s string) bool {
	if len(s) <= 2 {
		return true
	}
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]] += 1
		} else {
			m[s[i]] = 1
		}
	}
	ans := 0
	for _, v := range m {
		ans += v / 2 * 2
		if v%2 == 1 && ans%2 == 0 {
			ans++
		}
	}
	return ans == len(s)
}

//leetcode submit region end(Prohibit modification and deletion)
