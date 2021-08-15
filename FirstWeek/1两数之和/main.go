package main

import "fmt"

var nums = []int{2, 7, 11, 15}

func twoSum(nums []int, target int) []int {
	//使用哈希表，降低时间复杂度
	ns := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		//便利数组，用target减去nums[i]，如果这个数存在hash中，那么就存在，hash认为是O(1)的时间复杂度
		another := target - nums[i]
		if y, ok := ns[another]; ok {
			return []int{y, i}
		}
		//将数：索引存入hash表
		ns[nums[i]] = i

	}
	return []int{}

}

func main() {
	fmt.Println(twoSum(nums, 9))
}
