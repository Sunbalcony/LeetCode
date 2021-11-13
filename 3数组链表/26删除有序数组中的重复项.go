package main

import "fmt"

//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

/*
有序数组
删除重复项

*/

var nums26 = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//快慢指针
	slow := 0
	fast := 1
	for fast < len(nums) {
		if nums[slow] != nums[fast] { //1 升序数组，如果相邻的连个不相同 那么慢指针下一个元素就是它
			nums[slow+1] = nums[fast]
			slow++
		}
		fast++ //2 否则 slow下一个fast值相同就默认fast指针向后移动

	}
	fmt.Println(nums)
	return slow + 1

}

func main() {
	fmt.Println(removeDuplicates(nums26))

}
