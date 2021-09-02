package main

import "fmt"

/*
https://leetcode-cn.com/problems/climbing-stairs/
*/
func climbStairs(n int) int {
	/*	按照递归代码模板：
		1 终止条件
		2 当前层逻辑处理
		3 下探到下一次
		4 清扫当前状态 不用*/
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	ret := 0
	//结果
	b := 2
	//两步
	a := 1
	//一步
	for i := 3; i <= n; i++ {
		ret = a + b
		a = b
		b = ret
	}
	print(ret)
	/*
		具体分析思路：https://mp.weixin.qq.com/s/MDrjEhSnMnt8JblIsy-ZTg
		但是重复计算的部分还没搞清楚
	*/

}

func main() {
	aa := climbStairs(4)
	fmt.Println(aa)

}
