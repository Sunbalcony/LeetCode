package main

import "fmt"

//https://leetcode-cn.com/problems/search-insert-position/

/*
condition：
	一个排序数组和一个目标值
	必须使用时间复杂度为 O(log n) 的算法
	nums 为无重复元素的升序排列数组


*/

/*
	插入的位置如何处理？

*/

var nums = []int{1, 3, 5, 6, 9}

func searchInsert(nums []int, target int) int {

	//如果nums[right]小于target 那么直接插入到数组末尾就好了
	if nums[len(nums)-1] < target {
		return len(nums) //插入位置
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			//在mid+1~right
			left = mid + 1
		} else {
			//缩减范围将右边界设置为left~mid
			right = mid
		}

	}
	return right

}

func main() {
	fmt.Println(searchInsert(nums, 7))

}
