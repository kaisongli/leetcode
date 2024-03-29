package array

/**
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。

*/
func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}
	newLen := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[newLen] = nums[i]
			newLen++
		}
	}
	for i := newLen; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroes2(nums []int) {
	if len(nums) < 2 {
		return
	}
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
