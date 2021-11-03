package main

var nums = []int{1, 2, 3}

func subsets(nums []int) [][]int {
	var result [][]int
	for i := 0; i < len(nums); i++ {
		var newest []int
		for j := 0; j < len(result); j++ {
			newSet := result[][j] + nums[i]
			append(newest, newSet)

		}
		append(result, newest)

	}

}

func main() {
	subsets(nums)

}
