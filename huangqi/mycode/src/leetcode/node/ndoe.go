package node

/**
 * 反正链表
 */
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var next, pre *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

/**
 * 二叉树的前序遍历
 */
//1. 递归实现
func PreorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	preorderTraversal(root, &res)
	return res
}
func preorderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preorderTraversal(root.Left, res)
	preorderTraversal(root.Right, res)
}

//2. 使用辅助栈实现
func PreorderTraversalWithStack(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root} // 使用栈来辅助迭代
	res := make([]int, 0)

	for len(stack) > 0 {

	}
}
