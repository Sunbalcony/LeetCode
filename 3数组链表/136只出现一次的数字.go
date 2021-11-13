package main

//https://leetcode-cn.com/problems/single-number/

var nums136 = []int{4, 1, 2, 1, 2}

//自己写的复杂度太高。 时间:O(n) 空间O(n)
func singleNumber(nums []int) int {
	hashNum := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := hashNum[nums[i]]; ok {
			hashNum[nums[i]] += 1
		} else {
			hashNum[nums[i]] = 1
		}

	}
	for i, value := range hashNum {
		if value == 1 {
			return i

		}

	}
	return 0

}

func main() {
	singleNumber(nums136)

}
