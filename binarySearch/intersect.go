package binarySearch
//给定两个数组，编写一个函数来计算它们的交集。
//
//
//
// 示例 1：
//
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2]
//
//
// 示例 2：
//
// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[9,4]
//
//
//
// 说明：
//
//
// 输出结果中的每个元素一定是唯一的。
// 我们可以不考虑输出结果的顺序。
//
// Related Topics 排序 哈希表 双指针 二分查找
// 👍 287 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var res []int
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			if res == nil || nums1[i] > res[len(res) - 1]{
				res = append(res, nums1[i])
			}
			i ++
			j ++
		}else if nums1[i] < nums2[j]{
			i ++
		}else {
			j ++
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

