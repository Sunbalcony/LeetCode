package main

import "sort"

//https://leetcode-cn.com/problems/assign-cookies/
//题目大意：简单的说就是给孩子分饼干，分的饼干大小尽可能可以满足孩子的刚合适的胃口，让饼干最大程度的利用
func main() {

}

func findContentChildren(g []int, s []int) int {
	var result int
	//g  小孩胃口
	//s 饼干
	//先将两数组升序排列，然后将饼干满足每个小朋友也就是s中元素要大于g中元素
	sort.Ints(g)
	sort.Ints(s)
	gLen := len(g)
	sLen := len(s)
	i, j := 0, 0
	for i < gLen && j < sLen {
		if  g[i]<= s[j] {
			result += 1
			i += 1
			j += 1

		} else {
			j += 1

		}
	}
	return result

}
