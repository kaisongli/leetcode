package bd

/**
25. K 个一组翻转链表
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

进阶：

你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	cur, pre := head, dummy
	len := 0
	for head != nil {
		len++
		head = head.Next
	}
	//复位
	head = dummy.Next
	var temp *ListNode
	for i := 0; i < len/k; i++ {
		for j := 0; j < k-1; j++ {
			temp = cur.Next
			cur.Next = temp.Next
			temp.Next = pre.Next
			pre.Next = temp
		}
		pre = cur
		cur = pre.Next
	}
	return dummy.Next
}
