package array

/**
80. 删除有序数组中的重复项 II
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 最多出现两次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
示例 1：

输入：nums = [1,1,1,2,2,3]
输出：5, nums = [1,1,2,2,3]
解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。 不需要考虑数组中超出新长度后面的元素。

*/
func removeDuplicatesII(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	cur := 2
	for next := 2; next < n; next++ {
		if nums[cur-2] != nums[next] {
			nums[cur] = nums[next]
			cur++
		}
	}
	return cur
}
