package node

import (
	"math"
)

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

// 2. 使用辅助栈实现
func PreorderTraversalWithStack(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{root} // 使用栈来辅助迭代
	res := make([]int, 0)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, node.Val)

		// 先将右子节点入栈，再将左子节点入栈
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return res
}

/**
 * 二叉树的中遍历
 */
//1. 递归实现
//2. 使用辅助栈实现
func InorderTraversallWithStack(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)

	return res
}

/**
 * 二叉树的后序遍历
 */
//1. 递归实现
//2. 使用辅助栈实现
func PostorderTraversallWithStack(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	res := make([]int, 0)
	for len(stack) > 0 || root != nil {
		//左子节点不为空
		for root.Left != nil {
			//左子节点入栈
			stack = append(stack, root.Left)
			root = root.Left
		}
		//获取栈顶元素
		peek := stack[len(stack)]

		if peek.Right == nil {
			//弹出栈顶元素
			stack = stack[:len(stack)-1]
			res = append(res, peek.Val)
			root = peek
		} else {
			stack = append(stack, peek.Right)
		}
	}
	return res
}

// 617. 合并二叉树
func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// 如果两颗树都为空，返回空
	if root1 == nil && root2 == nil {
		return nil
	}
	// 如果第一颗树为空，返回第二颗树
	if root1 == nil {
		return root2
	}
	// 如果第二颗树为空，返回第一颗树
	if root2 == nil {
		return root1
	}
	// 创建新的树节点，值为两颗树节点值的和
	treeNode := &TreeNode{Val: root1.Val + root2.Val}
	// 递归合并左子树
	treeNode.Left = MergeTrees(root1.Left, root2.Left)
	// 递归合并右子树
	treeNode.Right = MergeTrees(root1.Right, root2.Right)
	return treeNode
}

// 226. 翻转二叉树
func InvertTree(root *TreeNode) *TreeNode {
	// 如果根节点为空，返回空
	if root == nil {
		return nil
	}
	// 创建镜像根节点，值与原根节点相同
	invert := &TreeNode{Val: root.Val}
	// 递归反转右子树，并赋给镜像根节点的左子树
	invert.Left = InvertTree(root.Right)
	// 递归反转左子树，并赋给镜像根节点的右子树
	invert.Right = InvertTree(root.Left)
	return invert
}

// 98. 验证二叉搜索树
func IsValidBST(root *TreeNode) bool {
	return isValidBST(root, math.MinInt64, math.MaxInt64)
}

// IsValidBST 是一个辅助函数，用于递归判断以给定节点为根的子树是否为二叉搜索树
// node: 当前节点
// min: 子树节点允许的最小值（不包含）
// max: 子树节点允许的最大值（不包含）
// 返回值:
// - 如果子树是二叉搜索树，则返回true
// - 如果子树不是二叉搜索树，则返回false
func isValidBST(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	// 如果节点的值大于等于最大值或小于等于最小值，则不满足二叉搜索树的定义
	if node.Val >= max || node.Val <= min {
		return false
	}

	// 递归判断左子树和右子树
	return isValidBST(node.Left, min, node.Val) && isValidBST(node.Right, node.Val, max)
}

// 二叉树的层序遍历
func LevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root} // 创建一个队列，初始时将根节点入队

	for len(queue) > 0 { // 当队列不为空时进行循环
		size := len(queue)         // 记录当前层的节点数量
		level := make([]int, size) // 创建一个临时切片，用于存储当前层的节点值

		for i := 0; i < size; i++ { // 遍历当前层的节点
			node := queue[i]    // 从队列中取出节点
			level[i] = node.Val // 记录节点值

			if node.Left != nil { // 如果节点存在左子节点，则将左子节点入队
				queue = append(queue, node.Left)
			}
			if node.Right != nil { // 如果节点存在右子节点，则将右子节点入队
				queue = append(queue, node.Right)
			}
		}

		res = append(res, level) // 将当前层的节点值切片添加到结果切片中
		queue = queue[size:]     // 移除已处理的节点，更新队列为下一层的节点
	}

	return res // 返回层序遍历的结果
}

/**
 * BM35 判断是不是完全二叉树
 *
 * @param root TreeNode类
 * @return bool布尔型
 */
func IsCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var queue []*TreeNode
	queue = append(queue, root)
	flag := false
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			flag = true
		} else {
			if flag {
				return false
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}

	}
	return true
}
