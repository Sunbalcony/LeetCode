package main

//https://leetcode-cn.com/problems/first-bad-version/

/*
每次版本为都是加1，可默认为一个递增数组
*/

func firstBadVersion(n int) int {
	left := 1
	right := n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			//isBadVersion(version)是Leetcode内函数
			//如果 isBadVersion(mid) 为true，那么意味着最少从此mid开始，往后版本都是错误的，因此答案区间可以缩减为left~mid
			right = mid //答案在区间 [left.mid]
		} else {
			left = mid + 1
		}

	}
	//此时有left==right,即为正确答案
	return left

}

func main() {
	firstBadVersion(7)

}
