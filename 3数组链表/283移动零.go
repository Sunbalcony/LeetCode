package main

import "fmt"

var nums = []int{0, 1, 0, 3, 12}

func moveZeros() []int {
	if len(nums) < 2 {
		return nums
	}
	j := 0
	//j指向下一个不为0的数字的存放位置
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	return nums

}

func main() {
	result := moveZeros()
	fmt.Println(result)
}
