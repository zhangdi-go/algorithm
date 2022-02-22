package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode // pre：前一个节点
	cur := head       // cur：当前的节点
	for cur != nil {
		tmp := cur.Next // 把下一个节点缓存起来
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
