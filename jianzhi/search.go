package jianzhi

/**
剑指 Offer 53 - I. 在排序数组中查找数字 I
统计一个数字在排序数组中出现的次数。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: 2
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: 0
*/

//排序数组首选二分查找，分别找到target的开始和结束位置，从而得出出现次数

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	start := findStartIndex(nums, target)
	if start == -1 {
		return 0
	}
	end := findEndIndex(nums, target)
	return end - start + 1
}

func findStartIndex(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left != len(nums) && nums[left] == target {
		return left
	}
	return -1
}

func findEndIndex(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right != -1 && nums[right] == target {
		return right
	}
	return -1
}
