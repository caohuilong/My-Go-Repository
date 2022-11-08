/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ll := &ListNode{}
	l := ll
	var c int
	for l1 != nil || l2 != nil || c != 0 {
		var k1, k2 int
		if l1 != nil {
			k1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			k2 = l2.Val
			l2 = l2.Next
		}
		tmp := &ListNode{
			Val: (c + k1 + k2) % 10,
		}
		c = (c + k1 + k2) / 10

		l.Next = tmp
		l = l.Next

	}
	return ll.Next
}

// @lc code=end

