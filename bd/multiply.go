package bd

//43 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
// 示例 1:
//
// 输入: num1 = "2", num2 = "3"
//输出: "6"
//
// 示例 2:
//
// 输入: num1 = "123", num2 = "456"
//输出: "56088"
//
// 说明：
//
//
// num1 和 num2 的长度小于110。
// num1 和 num2 只包含数字 0-9。
// num1 和 num2 均不以零开头，除非是数字 0 本身。
// 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
//
// Related Topics 数学 字符串
// 👍 595 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 || num1 == "0" || num2 == "0" {
		return "0"
	}
	n, m := len(num1), len(num2)
	ansArr := make([]int, m+n)
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			x, y := num1[i]-'0', num2[j]-'0'
			mul := x * y
			p1, p2 := i+j, i+j+1
			sum := int(mul) + ansArr[p2]
			ansArr[p2] = sum % 10
			ansArr[p1] += sum / 10
		}
	}
	i := 0
	for i < len(ansArr) && ansArr[i] == 0 {
		i++
	}
	ans := ""
	for ; i < len(ansArr); i++ {
		ans += string('0' + ansArr[i])
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
