package main

import "fmt"

//https://leetcode-cn.com/problems/binary-search/

/*
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target
题目给定了n个元素的升序数组，因此用二分查找比较合适
*/

var nums = []int{-1, 0, 3, 5, 9, 12}

func searchIndex(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			//如果是就返回
			return mid
		} else if nums[mid] > target {
			//如果中间大于target那么右边界就左移动
			right = mid - 1
		} else {
			//否则左边界就右移动
			left = mid + 1
		}
	}
	return -1

}

func main() {
	fmt.Println(searchIndex(nums, 3))
	fmt.Println(searchIndex(nums, -1))

}
