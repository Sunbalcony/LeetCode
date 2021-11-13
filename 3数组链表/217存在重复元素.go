package main

import "fmt"

//https://leetcode-cn.com/problems/contains-duplicate/

var nums217 = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

func containsDuplicate(nums []int) bool {
	temp := map[int]struct{}{}
	for i := 0; i < len(nums); i++ {
		if _, ok := temp[nums[i]]; ok {
			return true

		}
		temp[nums[i]] = struct{}{}
	}
	return false

}

func main() {
	fmt.Println(containsDuplicate(nums217))
	for i, i2 := range nums217 {
		fmt.Println(i, i2)

	}

}
