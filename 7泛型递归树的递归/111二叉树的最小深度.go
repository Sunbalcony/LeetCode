package main

//https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
叶子节点的定义是左孩子和右孩子都为 null 时叫做叶子节点
当 root 节点左右孩子都为空时，返回 1
当 root 节点左右孩子有一个为空时，返回不为空的孩子节点的深度
当 root 节点左右孩子都不为空时，返回左右孩子较小深度的节点值

*/

func minDepth1(root *TreeNode) int {
	//终止条件
	if root == nil {
		return 0
	}
	//condition one
	if root.Left == nil && root.Right == nil {
		return 1
	}

	depth1 := minDepth(root.Left) + 1
	depth2 := minDepth(root.Right) + 1
	//condition two
	if root.Left == nil || root.Right == nil { //叶子节点指的是左右孩子节点都为空的节点
		if depth1 > depth2 {
			return depth1 //这种情况下取其中最大的那个
		}
		return depth2

	}
	if depth1 > depth2 {
		return depth2
	}
	return depth1

}

//简化版本
func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	depth1 := minDepth(root.Left) + 1
	depth2 := minDepth(root.Right) + 1
	if root.Left == nil || root.Right == nil { //叶子节点指的是左右孩子节点都为空的节点
		if depth1 > depth2 {
			return depth1 //这种情况下取其中最大的那个
		}
		return depth2

	}
	if depth1 > depth2 {
		return depth2
	}
	return depth1

}

func main() {

}
