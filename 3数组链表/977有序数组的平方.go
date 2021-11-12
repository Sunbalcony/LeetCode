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

/*
最关键的点就是：原数组中的元素平方最大值一定产生在原数组的最左边或者最右边。
*/


//双指针时间复杂度O(n)
func sortedSquaresByDoublePointer(nums []int) []int {
	left := 0              //指向数组最左边
	right := len(nums) - 1 //指向数组最右边

	//创建一个新数组，存储平方值
	var newNums = make([]int,len(nums))
	write := len(nums) - 1 //得到元素平方值，从数组最后面开始写入
	for left <= right {    //左右指针相遇后遍不再循环
		//如果左指针数平方大于有指针数平方，就将左指针数平方放入数组最后一个位置
		if nums[left]*nums[left] > nums[right]*nums[right] {
			newNums[write] = nums[left] * nums[left]
			//左指针右移动
			left++
			//写数组指针往前移动一个位置
			write--
		} else {
			newNums[write] = nums[right] * nums[right]
			right--
			write--
		}

	}
	return newNums

}

func main() {
	fmt.Println("插入排序", sortedSquaresByInsertionSort(nums977))
	fmt.Println("冒泡排序", sortedSquaresByBubbleSort(nums977))
	fmt.Println("双指针", sortedSquaresByDoublePointer(nums977))

}
