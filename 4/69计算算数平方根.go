package main

//https://leetcode-cn.com/problems/sqrtx/

//二分法:为什么能用二分查找？
//y = x^2 (x>0),其实是这么一个函数
//抛物线在y轴右侧单调递增，且有上下界0,x 就可以用二分查找

func mySqrt(x int) int {
	if x==0 ||x==1{
		//平方根肯定落在0~1之间的
		return x
	}
	left:=1  //左边界
	right:=x //右边界

	mid:=1  //随便初始化一个值

	for left < right{
		mid = left+(right-left)/2

	}




}


func main() {

}
