package arrayandstring

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
// 示例 1:
//
// 输入: 123
//输出: 321
//
//
// 示例 2:
//
// 输入: -123
//输出: -321
//
//
// 示例 3:
//
// 输入: 120
//输出: 21
//
//
// 注意:
//
// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231, 231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
// Related Topics 数学
// 👍 2189 👎 0

/**
解法：
关键点：判断res是否超出int32的长度，将res转化为int32类型:int32(res)并赋值给temp，方便判断;
3.判断方法：提前赋值判断。
利用int32转换，判断tmp值，如果tmp是溢出的，那么(tmp*10)/10不会等于tmp，从而终止程序return 0;

*/
//leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
	if x == 0 {
		return x
	}
	rev := 0
	for x != 0 {
		if tmp := int32(rev); (tmp*10)/10 != tmp {
			return 0
		}
		pop := x % 10
		x /= 10
		rev = rev*10 + pop
	}
	return rev
}

//leetcode submit region end(Prohibit modification and deletion)
