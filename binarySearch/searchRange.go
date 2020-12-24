package binarySearch

//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
// 你的算法时间复杂度必须是 O(log n) 级别。
//
// 如果数组中不存在目标值，返回 [-1, -1]。
//
// 示例 1:
//
// 输入: nums = [5,7,7,8,8,10], target = 8
//输出: [3,4]
//
// 示例 2:
//
// 输入: nums = [5,7,7,8,8,10], target = 6
//输出: [-1,-1]
// Related Topics 数组 二分查找
// 👍 601 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	if nums == nil || len(nums) == 0{
		return []int{-1, -1}
	}
	firstPosition := findFirstPosition(nums, target)
	if firstPosition == -1{
		return []int{-1, -1}
	}
	lastPosition := findLastPosition(nums, target)
	return []int{firstPosition, lastPosition}
}

func findFirstPosition(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right{
		mid := left + (right - left) >> 1
		if nums[mid] == target {
			right = mid - 1
		}else if nums[mid] < target {
			left = mid + 1
		}else {
			right = mid - 1
		}
	}
	if left != len(nums) && nums[left] == target {
		return left
	}
	return - 1
}

func findLastPosition(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right{
		mid := left + (right - left) >> 1
		if nums[mid] == target {
			left = mid + 1
		}else if nums[mid] < target {
			left = mid + 1
		}else {
			right = mid - 1
		}
	}
	if right != -1 && nums[right] == target {
		return right
	}
	return -1
}


//leetcode submit region end(Prohibit modification and deletion)

