/*
 * @lc app=leetcode.cn id=725 lang=golang
 *
 * [725] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
	res := make([]*ListNode, k)
	len := 0
	for h := head; h != nil; h = h.Next {
		len++
	}

	m, n := len/k, len%k

	for i := 0; i < k && head != nil; i++ {
		res[i] = head
		if i == k-1 {
			break
		}
		partSize := m
		if i < n {
			partSize = m + 1
		}

		for j := 1; j < partSize; j++ {
			head = head.Next
		}
		head, head.Next = head.Next, nil
	}
	return res
}

// @lc code=end

