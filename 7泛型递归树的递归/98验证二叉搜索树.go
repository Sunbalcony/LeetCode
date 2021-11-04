package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//思路1：中序遍历时，判断当前节点是否大于中序遍历的前一个节点，如果大于，说明满足 BST，继续遍历；否则直接返回 false。
func isValidBST(root *TreeNode) bool {
	var stack []*TreeNode
	inorder := math.MinInt64 //给定一个初始值 记录上一个值。初始值为int64最小值 默认树中的最小值是大于等于int64的最小值
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root) //将每个节点压入栈中
			root = root.Left            //向左边走
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right

	}
	return true
	//复杂度分析：
	//时间复杂度：遍历了整个二叉树 时间复杂度为0n
	//空间复杂度：因为使用stack来装二叉树的每个节点，因此空间复杂度为0n

}

//中序遍历 递归版本
func isValidBSTRecursion(root *TreeNode) bool {
	var inorder func(root *TreeNode) bool
	tmp := math.MinInt64 //记录上一个值。初始值为int64最小值

	inorder = func(root *TreeNode) bool {
		//终止条件
		if root == nil {
			return true
		}
		l := inorder(root.Left)
		if root.Val <= tmp {
			return false
		}
		tmp = root.Val
		r := inorder(root.Right)
		return l && r

	}
	return inorder(root)

}

func main() {

}
