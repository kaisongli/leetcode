package array
//给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
//
//
//
// 说明：
//
//
// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
// 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
//
//
//
//
// 示例：
//
//
//输入：
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//输出：[1,2,2,3,5,6]
//
//
//
// 提示：
//
//
// -10^9 <= nums1[i], nums2[i] <= 10^9
// nums1.length == m + n
// nums2.length == n
//
// Related Topics 数组 双指针
// 👍 707 👎 0

//三指针+尾插法
//leetcode submit region begin(Prohibit modification and deletion)
func merge(nums1 []int, m int, nums2 []int, n int)  {
	p, q, k := m-1, n-1, m+n-1
	for p >= 0 && q >= 0 {
		if nums1[p] < nums2[q] {
			nums1[k] = nums2[q]
			q --
		}else {
			nums1[k] = nums1[p]
			p --
		}
		k --
	}
	for q >= 0 {
		nums1[k] = nums2[q]
		q --
		k --
	}
}
//leetcode submit region end(Prohibit modification and deletion)

