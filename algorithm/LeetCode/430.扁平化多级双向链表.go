/*
 * @lc app=leetcode.cn id=430 lang=golang
 *
 * [430] 扁平化多级双向链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	flattenChild(root)
	return root
}

func flattenChild(node *Node) (last *Node) {
	cur := node

	for cur != nil {
		next := cur.Next
		if cur.Child != nil {
			tmp := flattenChild(cur.Child)
			cur.Next = cur.Child
			cur.Child = nil
			cur.Next.Prev = cur
			tmp.Next = next
			if next != nil {
				next.Prev = tmp
			}
			last = tmp
		} else {
			last = cur
		}
		cur = next
	}
	return last
}

// @lc code=end

