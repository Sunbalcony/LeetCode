package main

import "fmt"

//https://leetcode-cn.com/problems/rotate-array/

var nums = []int{1, 2, 3, 4, 5, 6, 7}

//func rotate(nums []int, k int) {
//	nums1 := nums[:k+1]
//	nums2 := nums[k+1:]
//	nums2 = append(nums2, nums1...)
//	fmt.Println(nums2)
//	nums = nums2
//
//}

func rotate(nums []int, k int) {
	k %= len(nums) //k可能比数组大，在这种情况下，向右移动整个数组长度后就回到原来的位置，直接%掉就是去掉每次回到原来位置的步数，获得最后移动的步数
	// 对数组长度进行取模，如果旋转13次，数组长度为4，那么13对4取模，剩1，只需一次就好了。 目的是为了减少旋转次数
	reserve(nums)
	reserve(nums[:k])
	reserve(nums[k:])
	fmt.Println(nums)

}

func reserve(a []int) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		//从数组中间切分i<n/2，翻转前后部分n-1-i是数组长度为n，-1是计算索引位，-i是a[i]对应到数组后半段从末尾向中间逼近一一对应的位置
		a[i], a[n-1-i] = a[n-1-i], a[i]

	}
}

func rotate2(nums []int, k int) {
	newNums := make([]int, len(nums))
	copy(newNums, nums)
	for i := 0; i < len(nums); i++ {
		nums[(i+k)%len(nums)] = newNums[i]
	}

}

func main() {
	rotate(nums, 3)

}
