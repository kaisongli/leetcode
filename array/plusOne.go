package array
/**
第66题：加一
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。
 */
func plusOne(digits []int) []int {
	addon := 0
	for i := len(digits) - 1; i >= 0; i -- {
		digits[i] += addon
		//复位
		addon = 0
		//如果是最后一位, 直接加1
		if i == len(digits)  - 1{
			digits[i] ++
		}
		//如果有进位
		if digits[i] == 10 {
			digits[i] = digits[i] % 10
			addon = 1
		}
	}
	var res []int
	//如果第一个数有进位
	if addon == 1{
		res = make([]int, 1)
		res[0] = 1
		res = append(res, digits...)
	}else {
		res = digits
	}
	return res
}
