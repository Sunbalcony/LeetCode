package main

// TreeNode https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

// N个遍历树的问题 https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/solution/yi-tao-quan-fa-shua-diao-nge-bian-li-shu-de-wen--3/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InorderTraversal1 迭代法
func InorderTraversal1(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode //手动维护一个stack
	for root != nil || len(stack) > 0 {
		for root != nil { //边界条件是root=！空，
			stack = append(stack, root) // 1：向栈中压入节点，每次都向左边走，知道到达边界条件root为nil跳出
			root = root.Left
		}
		res = append(res, root.Val)//2：将该节点的val添加到结果集中
		root = stack[len(stack)-1]   // 3：回到上一节点

		stack = stack[:len(stack)-1] //4： 将该节点出栈

		root = root.Right //5：检查向右叶子节点存不存在

	}
	return res

}

// InorderTraversal2 递归法
func InorderTraversal2(root *TreeNode) []int {
	var res []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right) //加入4已经处理过了到2，进入递归节点不为空，下一步去left，为空return后，记录当前2节点的Val，然后进入right，right不为空，但是right下面无left，那么return，记录当前right节点的值即可

	}
	inorder(root)
	return res
	/*
	   		1
	         /   \
	       2       3
	     /  \      /
	   4      5   6
	*/

}

func main() {

}
