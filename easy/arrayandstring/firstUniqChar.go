package arrayandstring
//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
//
//
//
// 示例：
//
// s = "leetcode"
//返回 0
//
//s = "loveleetcode"
//返回 2
//
//
//
//
// 提示：你可以假定该字符串只包含小写字母。
// Related Topics 哈希表 字符串
// 👍 267 👎 0

/**
解法：
1、定义字符数组，26长度
2、先给字符赋值为对应下标
3、再次遍历时，如果字符位置等于对应的下标，说明没有重复字符，返回即可，否则字符位置重新赋值为-1
 */

//leetcode submit region begin(Prohibit modification and deletion)
func firstUniqChar(s string) int {
	if len(s) < 1 {
		return -1
	}
	if len(s) == 1 {
		return 0
	}
	var chars [26]int
	for i, ch := range s{
		chars[ch - 'a'] = i
	}
	for i, ch := range s {
		if chars[ch - 'a'] == i {
			return i
		}else {
			chars[ch - 'a'] = -1
		}
	}
	return -1
}
//leetcode submit region end(Prohibit modification and deletion)
