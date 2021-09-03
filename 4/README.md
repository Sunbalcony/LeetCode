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

#贪心算法与动态规划

##贪心算法

贪心法难就难在你怎么证明他就是可以用贪心法就可以进行解决 

适用场景

简单的说，问题能够分解成子问题来解决，子问题的最优解能递推到最终问题的最优解。这种子问题最优解称为最优子结构

贪心算法与动态规划的不同在于它对每个子问题的解决方案都做出选择，不能<b>回退<b>。

动态规划则会保存之前的运算结果，并根据以前的结果对当前进行选择，有回退功能