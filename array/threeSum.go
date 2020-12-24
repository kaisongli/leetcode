package array

import "sort"

/**
第15题：三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。注意：答案中不可以包含重复的三元组。
示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
 */

func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i ++ {
		if nums[i] > 0{
			return res
		}
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left] + nums[right] + nums[i] == 0 {
				res = append(res, []int{nums[left], nums[right], nums[i]})
				//去重
				for left < right && nums[left] == nums[left + 1] {
					left ++
				}
				for left < right && nums[right] == nums[right - 1]{
					right --
				}
				left ++
				right --
			}else if nums[left] + nums[right] + nums[i] > 0 {
				right --
			}else {
				left ++
			}
		}
	}
	return res
}
