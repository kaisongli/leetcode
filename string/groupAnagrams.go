package string

import (
	"sort"
	"strings"
)

//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例:
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//
// 说明：
//
//
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。
//
// Related Topics 哈希表 字符串
// 👍 537 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	ans := make([][]string, 0)
	m := map[string]int{}
	index := 0
	for _, str := range strs {
		chars := strings.Split(str, "")
		sort.Strings(chars)
		toKey := strings.Join(chars, "")
		if idx, ok := m[toKey]; ok {
			ans[idx] = append(ans[idx], str)
		}else {
			m[toKey] = index
			index ++
			ans = append(ans, []string{str})
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
