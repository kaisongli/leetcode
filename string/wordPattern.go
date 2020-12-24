package string

import "strings"

//给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。
//
// 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。
//
// 示例1:
//
// 输入: pattern = "abba", str = "dog cat cat dog"
//输出: true
//
// 示例 2:
//
// 输入:pattern = "abba", str = "dog cat cat fish"
//输出: false
//
// 示例 3:
//
// 输入: pattern = "aaaa", str = "dog cat cat dog"
//输出: false
//
// 示例 4:
//
// 输入: pattern = "abba", str = "dog dog dog dog"
//输出: false
//
// 说明:
//你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。
// Related Topics 哈希表
// 👍 217 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s, " ")
	if len(strs) != len(pattern) {
		return false
	}
	patternMap := map[byte]string{}
	sMap := map[string]byte{}
	for i := 0; i < len(pattern); i ++ {
		_, patternOk := patternMap[pattern[i]]
		_, sOk := sMap[strs[i]]
		if (patternOk && patternMap[pattern[i]] != strs[i]) || (sOk && sMap[strs[i]] != pattern[i]){
			return false
		}else {
			patternMap[pattern[i]] = strs[i]
			sMap[strs[i]] = pattern[i]
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
