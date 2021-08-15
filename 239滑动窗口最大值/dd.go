package main

import "fmt"

var a2 = []int{1, 3, -1, -3, 5, 3, 6, 7}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	window := []int{}
	l, r := 0, 0
	result := []int{}
	for ; r < n; r++ {
		if len(window) == 0 {
			window = append(window, r)

		} else if nums[r] > nums[window[len(window)-1]] {
			for nums[r] > nums[window[len(window)-1]] {
				window = window[:len(window)-1]
				if len(window) == 0 {
					break
				}

			}
			window = append(window, r)
		} else {
			window = append(window, r)
		}
		if r>k-1{
			l++
		}
	}
	fmt.Println(result)
	return result

}

func main() {
	maxSlidingWindow(a2 ,3)

}
