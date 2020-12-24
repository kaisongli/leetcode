package arrayandstring
//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
//示例 1:
//
//输入: s = "anagram", t = "nagaram"
//输出: true
//
//
// 示例 2:
//
// 输入: s = "rat", t = "car"
//输出: false
//
// 说明:
//你可以假设字符串只包含小写字母。
//
// 进阶:
//如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
// Related Topics 排序 哈希表
// 👍 249 👎 0

/**
解法：
先判断长度，如果长度不相等，直接返回false
定义长度26的字母数组，s负责+1，t负责-1
最终遍历字母数组，如果不等于0，返回false
 */
//leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var chars [26]int
	for i, ch := range s{
		chars[ch - 'a'] ++
		chars[t[i] - 'a'] --
	}
	for _, v := range chars {
		if v != 0 {
			return false
		}
	}
	return true
}
