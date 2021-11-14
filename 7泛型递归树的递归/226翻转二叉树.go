package main

//https://leetcode-cn.com/problems/invert-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归1：先将两个子树进行翻转，再翻转子树内部
//先交换左右子树 再翻转
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left //交换二叉树左右子树

	//翻转子树内部
	invertTree1(root.Left)
	invertTree1(root.Right)
	return root

}

//递归2：再翻转子树内部，再交换左右子树
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	//翻转子树内部
	invertTree2(root.Left)
	invertTree2(root.Right)
	root.Left, root.Right = root.Right, root.Left //交换二叉树左右子树
	return root

}

func main() {
	invertTree2(&TreeNode{})
	invertTree1(&TreeNode{})

}
