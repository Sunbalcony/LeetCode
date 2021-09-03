package main

//https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

//既然是按层次遍历，那便是广度优先搜索BFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int

	if root == nil {
		return result
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		result = append(result, []int{})
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			node := q[j]
			result[i] = append(result[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}

		}
		q = p

	}
	return result

}

func main() {
	//时间复杂度，每个节点各进出队列一次 因此为O(n)
	//优秀讲解https://leetcode-cn.com/problems/binary-tree-level-order-traversal/solution/bfs-de-shi-yong-chang-jing-zong-jie-ceng-xu-bian-l/

}
