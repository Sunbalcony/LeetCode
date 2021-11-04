package main

//https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth1 := 1
	depth2 := 1
	depth1 += maxDepth(root.Left)
	depth2 += maxDepth(root.Right)
	if depth1 > depth2 {
		return depth1
	}
	return depth2

}

func main() {
	maxDepth(&TreeNode{})

}
