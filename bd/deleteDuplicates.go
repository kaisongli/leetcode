package bd

//存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。
//
// 返回同样按升序排列的结果链表。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,3,4,4,5]
//输出：[1,2,5]
//
//
// 示例 2：
//
//
//输入：head = [1,1,1,2,3]
//输出：[2,3]
//
//
//
//
// 提示：
//
//
// 链表中节点数目在范围 [0, 300] 内
// -100 <= Node.val <= 100
// 题目数据保证链表已经按升序排列
//
// Related Topics 链表
// 👍 585 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//哨兵节点
	dummy := &ListNode{Next: head}
	first, sec := dummy, dummy.Next
	for sec != nil && sec.Next != nil {
		if sec.Val == sec.Next.Val {
			for sec.Next != nil && sec.Val == sec.Next.Val {
				sec = sec.Next
			}
			first.Next = sec.Next
		} else {
			first = first.Next
		}
		sec = sec.Next
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
