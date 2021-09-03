package main

// Node https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var result []int
	var prevorder func(node *Node)
	prevorder = func(node *Node) {
		if node==nil{
			return
		}
		result = append(result,node.Val)//front
		for i := 0; i < len(node.Children); i++ { //遍历N叉树，N叉树是一个数组，因此只需要遍历一下整个数组即可
			prevorder(node.Children[i])


		}

	}
	prevorder(root)
	return result


}

func main() {

}
