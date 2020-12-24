package string

/**
第387题：字符串中的第一个唯一字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1 。
案例:

s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.
注意事项： 您可以假定该字符串只包含小写字母。

解法：
由于字母共有 26 个，所以我们可以声明一个 26 个长度的数组（该种方法在本类题型很常用）
因为字符串中字母可能是重复的，所以我们可以先进行第一次遍历，在数组中记录每个字母的最后一次出现的所在索引。
然后再通过一次循环，比较各个字母第一次出现的索引是否为最后一次的索引。
如果是，我们就找到了我们的目标，如果不是我们将其设为 -1（标示该元素非目标元素）如果第二次遍历最终没有找到目标，直接返回 -1即可。
 */
func firstUniqChar(s string) int {
	if len(s) < 1 {
		return -1
	}
	if len(s) == 1 {
		return 0
	}
	var arr [26]int
	for i, ch := range s{
		arr[ch - 'a'] = i
	}
	for i, ch := range s{
		if arr[ch - 'a'] == i {
			return i
		}else {
			arr[ch - 'a'] = -1
		}
	}
	return -1
}
