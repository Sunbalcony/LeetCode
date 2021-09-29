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

 
## 二分查找的前提

### 1 目标函数存在单调性（单调递增或递减）

二分查找指明的是在有序的里面进行查找，如果无序，是无法进行二分查找的，无序只能从头到尾遍历。
正因为是有序的，因此可以判断某些特征排除掉前半段或者后半段

### 2 存在上下界

没有上下界 空间就是无穷大了

### 3 能够通过索引访问


###代码模板

```markdown
left,right = 0,len(array) -1
左右边界分别是0和数组长度减1，也就是左右下标值
while left<=right:
    mid = (left + right)/2
    if array[mid]==target:
    判断mid是否等于target
        break or return result
    elif array[mid] < target:
        left = mid +1
    else:
        right = mid +1

```
递增数组,不连续,整体单调递增,查找31

[10,14,19,26,27,31,33,35,42,44]

如果普通查找方法，就是遍历，第六次可以查找出来

### 二分

index_start = 0

index_end = len(array)-1


索引出来后队列的数值 27

31 > 27

可以排除前半部分

继续从后半部分查找

index_start = (len(array)-1) /2

index_end = len(array)-1

索引出来的数字 35

35 > 31

。。。。。。。












