package main

//https://leetcode-cn.com/problems/sqrtx/

//给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
//由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

//为什么可以用二分查找

// y =x^2 x>0 在y轴右侧单调递增 上下界: 0，x

//二分查找的下界为 00，上界可以粗略地设定为 x。
//在二分查找的每一步中，我们只需要比较中间元素 mid 的平方与 xx 的大小关系，
//并通过比较的结果调整上下界的范围。
//由于我们所有的运算都是整数运算，不会存在误差，因此在得到最终的答案 ans 后，也就不需要再去尝试 ans+1 了。

func mySqrt2(x int) int {
	//超时
	if x == 0 || x == 1 {
		return x

	}
	left := 0
	right := x //right设置为x的平方根肯定是落在1和x之间的
	mid := 1   //随意一个
	for left <= right {
		mid = left + (right-left)/2
		if mid*mid < x {
			left = mid + 1
		} else if mid*mid > x {
			right = mid - 1
		}
	}
	return right

}

func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

func main() {
	mySqrt(1)
}
