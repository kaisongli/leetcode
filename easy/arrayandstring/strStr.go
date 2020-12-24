package arrayandstring
//实现 strStr() 函数。
//
// 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如
//果不存在，则返回 -1。
//
// 示例 1:
//
// 输入: haystack = "hello", needle = "ll"
//输出: 2
//
//
// 示例 2:
//
// 输入: haystack = "aaaaa", needle = "bba"
//输出: -1
//
//
// 说明:
//
// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//
// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
// Related Topics 双指针 字符串
// 👍 575 👎 0
/**
解法：
1、先判断needle字符串长度，如果大于haystack，肯定不匹配，返回-1
2、如果needle字符串长度等于0 ，直接return 0
3、分别获取haystack 和 needle的字符串长度， 遍历haystack字符串
如果start + needleLen长度的字符串等于needle， 返回start
4、否则返回-1
*/

//leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	if len(needle) == 0 {
		return 0
	}
	haystackLen, needleLen := len(haystack), len(needle)
	for start := range haystack[:(haystackLen - needleLen + 1)]{
		if haystack[start: (start + needleLen)] == needle {
			return start
		}
	}
	return -1
}
//leetcode submit region end(Prohibit modification and deletion)

