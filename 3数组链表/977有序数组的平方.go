package main

import "fmt"

//https://leetcode-cn.com/problems/squares-of-a-sorted-array/

var nums977 = []int{-7, -3, 2, 3, 11}

func sortedSquaresByInsertionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	//fmt.Println(nums)

	//插入排序
	if len(nums) <= 1 {
		return nums
	}
	for i := 1; i < len(nums); i++ {
		//index=0作为已排序区间元素，所以从index=1开始。初始已排序区间只有一个元素，那就是数组的第一个元素
		//插入排序的核心思想是取未排序区间元素在已排序区间中找到合适的位置进行插入，并保证已排序区间一直有序，并重复这个过程
		value := nums[i]
		j := i - 1
		//查找插入的位置
		for ; j >= 0; j-- {
			if nums[j] > value { //如果现在排过序的元素大于目标元素
				nums[j+1] = nums[j] //数据移动
			} else {
				break
			}

		}
		nums[j+1] = value //插入数据
	}
	return nums

}

func sortedSquaresByBubbleSort(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}

	//冒泡排序
	if len(nums) <= 1 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]

			}

		}

	}
	return nums
}

func main() {
	fmt.Println(sortedSquaresByInsertionSort(nums977))
	fmt.Println(sortedSquaresByBubbleSort(nums977))
	
}
