package main

//https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/

/*
给定的是非递减顺序序列的整数数组
*/

var nums = []int{2, 7, 11, 15}

/*
二分查找法，可以先固定一个数，然后寻找第二个数，如果第二个数等于目标数减去第一个的差
*/
func twoSumByBinarySearch(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for i := 0; i < len(numbers); i++ {

	}

}

/*
双指针解法

双指针解法本质：每次可以缩减一行或者一列的搜索结果
https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/solution/yi-zhang-tu-gao-su-ni-on-de-shuang-zhi-zhen-jie-fa/
*/

func twoSumByDoublePointer(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right { //非递减有序数组 在left<right返回遍历，本质是左右边界互相逼近
		if numbers[left]+numbers[right] == target { //如果两数相加相等就返回
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] < target { //如果left + right 小于 target 那说明左边界大小不够大，将左边界右移
			left++

		} else { //反之右边界左移动
			right--
		}
	}
	return []int{-1, -1}

	//时间复杂度 0(n)
	//空间复杂度 0(n)

}

func twoSumByDoublePointerSecond(numbers []int, target int) []int {
	//某种情况下来说应该让 ==target的这个式子放在else里面，效率会更高
	left := 0
	right := len(numbers) - 1
	for left < right { //非递减有序数组 在left<right返回遍历，本质是左右边界互相逼近
		if numbers[left]+numbers[right] > target { //如果两数相加相等就返回
			right--
		} else if numbers[left]+numbers[right] < target { //如果left + right 小于 target 那说明左边界大小不够大，将左边界右移
			left++
		} else { //反之右边界左移动
			return []int{left + 1, right + 1}
		}
	}
	return []int{-1, -1}

	//时间复杂度 0(n)
	//空间复杂度 0(n)

}

func main() {

}
