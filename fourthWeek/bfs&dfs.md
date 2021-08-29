# DFS与BFS

先看看在二叉树上进行DFS遍历和BFS遍历

DFS遍历使用 <b color=blue>递归<b>

```go

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	dfs(root.Right)
}
```

BFS遍历使用 <b color=blue>队列<b> 数据结构

```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(root *TreeNode) {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}

}
```