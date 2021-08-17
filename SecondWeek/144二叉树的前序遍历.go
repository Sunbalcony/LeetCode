package main

//https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

type TreeNodeF struct {
	Val   int
	Left  *TreeNodeF
	Right *TreeNodeF
}

//递归法
func preorderTraversal(root *TreeNodeF) []int {
	var result []int
	var preorder func(node *TreeNodeF)
	preorder = func(node *TreeNodeF) {
		if node == nil { //这个条件每次老忘，一定要记得
			return
		}
		result = append(result, node.Val) //根节点
		preorder(node.Left)  //左子树
		preorder(node.Right) //右子树
	}
	preorder(root)
	return result

}

//迭代法
func preorderTraversal1(root *TreeNodeF) []int {
	var result []int
	var stack []*TreeNodeF
	for root != nil || len(stack) > 0 {
		for root != nil {
			result = append(result,root.Val) //根：到那个节点 哪个节点的值先入result数组
			stack = append(stack,root) //将节点压入stack
			root = root.Left //左子树遍历
		}
		//如果root.left为空
		root = stack[len(stack)-1] //那么就回到上一个结果
		stack = stack[:len(stack)-1]//出栈
		root = root.Right //看看右边有没有

	}
	return result

}

func main() {

}
