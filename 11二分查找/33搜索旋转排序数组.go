package main

/*
题目要求 O(logN) 的时间复杂度，基本可以断定本题是需要使用二分查找，怎么分是关键
33 搜索旋转排序数组
切题，没思路
1 暴力求解 遍历数组 时间复杂度O(n) 还原  升序 ，然后二分 （关键是还原是怎么还原  最差还是O(n),可以O(logN也是用二分去找的)）


正解：二分查找
 - 单调
 - 边界
 - 可以index访问

我们根据nums[mid]和target的大小来判断target是在mid的左边还是右边 从而来调节左右边界left和right
但是对于旋转数组 我们无法直接根据mid和target的关系来判断target在mid的左侧还是右侧，因此需要“分段讨论”

*/

func search(nums []int, target int) int {
	left := 0              //左边界
	right := len(nums) - 1 //右边界
	for left < right {
		mid := left + (right-left)/2
		if mid == target {
			return mid
		}
		//根据nums[mid]和nums[left]的关系判断mid在左边还是右边
		if nums[mid] > nums[left] {
			//再判断target是在mid的左边还是右边，从而调整边界left和right
			if target > nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}

		} else {
			if target > nums[mid] && target < nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}

		}
	}
	return -1
}

func main() {

}
