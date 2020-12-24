package array

import "strings"

/**
题目14: 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，则返回""
示例1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释:

输入不存在公共前缀。
说明：

所有输入只包含小写字母 a-z
 */

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pref := strs[0]
	for i := 1; i < len(strs); i ++ {
		for strings.Index(strs[i], pref) != 0 {
			if len(pref) == 0 {
				return ""
			}
			pref = pref[: len(pref) - 1]
		}
	}
	return pref
}