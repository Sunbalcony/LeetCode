package main

/*
https://leetcode-cn.com/problems/generate-parentheses/
*/

//按照视频里的思路，复现整个写递归代码的过程

//数字 n 代表生成括号的对数，一对括号由左右括号合并组成，因此一共有2*n的格子需要我们填满括号那么我们先来填满括号

func addParenthesis(level int, max int, s string) {
	/*	0是格子的起始位置
		max是格子的最大位置
		""是当前字符串的样子
	*/
	//按照递归模板
	//终止条件
	if level >= max {
		println(s)
		return
	}
	//处理当前逻辑
	//想象在这一层 给进来的s添加左括号或者是右括号
	s1 := s + "("
	s2 := s + ")"
	//进入下一层
	addParenthesis(level+1, max, s1)
	addParenthesis(level+1, max, s2)
	//清理状态，不需要

}

/*
	上面函数遍历了2对括号所有可能的结果
	((((
	((()
	(()(
	(())
	()((
	()()
	())(
	()))
	)(((
	)(()
	)()(
	)())
	))((
	))()
	)))(
	))))
我们要在里面找到成对合法的符合要求的括号组合
如何判断括号合不合法？那就是在当前逻辑层进行判断
	//一个括号如何判断合不合法？
	//这里有两点：
	//左边括号随时可以增加，因为是左括号，只要它的数量不超过n(n对括号：n个左括号，n个右括号)即可
	//右括号一开始出现是不行的，一直出直接就是不合法的，右括号必须出现在左括号的后面，所以加右括号必须之前有左括号垫背。并且，左括号的数量要大于右括号
看下面的函数
*/
//2^n
func generateParenthesis(left, right, n int, s string) []string {
	//这里我们就不按照之前的数格子，而是要记用了多少左括号和右括号
	//按照递归模板
	//终止条件
	var result []string
	if left == n && right == n { //左和右都用完了
		println(s)
		result = append(result, s)
		return result

	}

	//处理当前逻辑
	if left < n { //如果左括号还没用完
		s1 := s + "("
		//进入下一层
		generateParenthesis(left+1, right, n, s1)

	}
	if left > right { //如果左括号比右括号多
		s2 := s + ")"
		//进入下一层
		generateParenthesis(left, right+1, n, s2)

	}

	//清理状态，不需要
	return nil

}

func main() {
	//addParenthesis(0, 4, "")
	//0是格子的起始位置
	//max是格子的最大位置
	//""是当前字符串的样子
	generateParenthesis(0, 0, 3, "")

}
