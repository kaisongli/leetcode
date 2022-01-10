package node

/**
第19题：删除链表倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。
进阶：

你能尝试使用一趟扫描实现吗？
*/
//哨兵节点+双指针
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n < 1 {
		return head
	}
	//哨兵节点
	pre := &ListNode{Next: head}
	//双指针
	first := pre
	sec := pre
	//先走n+1步
	for i := 0; i < n+1; i++ {
		if first != nil {
			first = first.Next
		} else {
			return nil
		}
	}
	//走到前置节点
	for first != nil {
		first = first.Next
		sec = sec.Next
	}
	//删除节点
	sec.Next = sec.Next.Next
	return pre.Next
}
