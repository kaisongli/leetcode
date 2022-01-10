package binarySearch

/**
74. 搜索二维矩阵
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。

示例 1：
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
输出：true
*/
//解法：二维数组数据拼接成一维数组，一次递增的，转化成普通的二分查找

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := left + (right-left)>>1
		x := matrix[mid/n][mid%n]
		if x == target {
			return true
		} else if x < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
