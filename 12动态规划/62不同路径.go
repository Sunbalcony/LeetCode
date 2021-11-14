package main

///https://leetcode-cn.com/problems/unique-paths/

/*
m x n的网格
机器人只能每次向下或者向右一步
*/

/*
https://leetcode-cn.com/problems/unique-paths/solution/san-chong-shi-xian-xiang-xi-tu-jie-62-bu-4jz1/
分析：
	每次只能向右或者向下一步，也就意味着机器人走路演化成了两个子问题
	当机器人真的向下或者向右走了一步，那么再下一步也是要考虑向下走还是向右走
	于是 该问题演化成了多个子问题
	我们只要将机器人从起点开始向下河向右方向所有不同种的走法加起来
	就是我们最终的解法
	path(start,end) = path(down,end) + path(right,end)
                        |                |
                        |                =path(right_down,end)+path(right_right,end)
                        |
                        =path(down_down,end)+path(down_right,end)
	递归：
	那么从(0,0)出发到(m,n)的所有路径是由两条线路加起来的
	所以递归的核心逻辑就是
	result = dfs(i+1,j)+dfs(i,j+1)
	有点类似于爬楼梯
	递推公式写出来了，那么终止条件是什么呢？
	终止条件就是当i或j到达网格边界会出发递归终止返回
	也就是当移动到最后面一列或者最下面一行
	也就是当i==m-1或j==n-1


	动态规划：
	对于(0,0)这个点来说，他只能往下走或者往右走
	那么反过来那一个点可以到达(2,2)呢？
		只能是他上方的(1,2)
		或者是他左边的(2,1)
	所以
	dp[i][j] = dp[i-1][j] + dp[i][j-1]
	核心逻辑有了，再处理边界条件
	【递归的边界条件】是走到网格最右边一列或者最下面一行
	【动态规划】正好是反过来的，因此我们要处理第一行和第一列，将他们都赋予1即可



*/

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)

	//第一行都赋予1
	for i := 0; i < m; i++ {
		dp[i] = make([]int,n)
		dp[i][0] = 1
	}
	//第一列都赋予1
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	//两个for循环推导，对于(i,j)来说，只能由上方或者左方转移而来
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]

}

func main() {

}
