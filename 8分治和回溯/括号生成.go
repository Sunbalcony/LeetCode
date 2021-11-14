package main

import "fmt"

//数字n代表生成括号的对数 ，一个括号由左右两个括号构成，故一共需要2n个括号 一共2^n个格子可以被填充

func addParenthesis(level int, max int, s string) {
	//终止条件
	if level > max {
		fmt.Println(s)
		return
	}
	//处理当前层逻辑
	str1 := s + "("
	str2 := s + ")"

	//进入下一层
	addParenthesis(level+1, max, str1)
	addParenthesis(level+1, max, str2)

}

func generateParenthesis(left, right int, n int, s string) {
	//终止条件
	var result []string
	if left == n && right == n {
		fmt.Println(s)
		result = append(result, s)

	}

	//当前层逻辑
	if left < n {
		str1 := s + "("
		generateParenthesis(left+1, right, n, str1)
	}
	if left > right {
		str2 := s + ")"
		generateParenthesis(left, right+1, n, str2)
	}

}

func main() {
	//addParenthesis(0, 8, "")
	generateParenthesis(0, 0, 4, "")

}
