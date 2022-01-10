package node

/**
编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。

示例1:

 输入：[1, 2, 3, 3, 2, 1]
 输出：[1, 2, 3]
示例2:

 输入：[1, 1, 1, 1, 2]
 输出：[1, 2]
*/
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	m := map[int]bool{
		head.Val: true,
	}
	cur := head
	for cur.Next != nil {
		tmp := cur.Next
		if !m[tmp.Val] {
			cur = cur.Next
			m[tmp.Val] = true
		} else {
			cur.Next = cur.Next.Next
		}
	}
	return head
}
