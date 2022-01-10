package binarySearch

import "time"

//实现 int sqrt(int x) 函数。
//
// 计算并返回 x 的平方根，其中 x 是非负整数。
//
// 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
//
// 示例 1:
//
// 输入: 4
//输出: 2
//
//
// 示例 2:
//
// 输入: 8
//输出: 2
//说明: 8 的平方根是 2.82842...,
//     由于返回类型是整数，小数部分将被舍去。
//
// Related Topics 数学 二分查找
// 👍 508 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func mySqrt(x int) int {
	if x == 1{
		return x
	}
	left, right := 1, x / 2
	for left <= right {
		mid := left + (right - left) >> 1
		if mid * mid == x {
			return mid
		}else if mid * mid > x {
			right = mid - 1
		}else{
			left = mid + 1
		}
	}
	return right
}
//leetcode submit region end(Prohibit modification and deletion)




2、select T1.name from T1 left join T2 on T1.id = T2.id where T2.score > 80 group by T1.name