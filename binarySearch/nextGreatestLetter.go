package binarySearch

/**
给你一个排序后的字符列表 letters ，列表中只包含小写英文字母。另给出一个目标字母 target，请你寻找在这一有序列表里比目标字母大的最小字母。

在比较时，字母是依序循环出现的。举个例子：

如果目标字母 target = 'z' 并且字符列表为 letters = ['a', 'b']，则答案返回 'a'


示例：

输入:
letters = ["c", "f", "j"]
target = "a"
输出: "c"

输入:
letters = ["c", "f", "j"]
target = "c"
输出: "f"

输入:
letters = ["c", "f", "j"]
target = "d"
输出: "f"
*/
func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1
	for left <= right {
		mid := left + (right-left)>>1
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return letters[left%len(letters)]
}
