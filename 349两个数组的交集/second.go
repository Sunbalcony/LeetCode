package main

import (
	"fmt"
	"sort"
)

var nums1 = []int{3, 4, 5, 6}
var nums2 = []int{7, 6, 5, 4, 3, 2}

func intersection() (inter []int) {
	//数组类问题，思考双指针，用双指针的前提，大体上是先对数组进行排序
	//快速排序O(n*logn)
	sort.Ints(nums1)
	sort.Ints(nums2)
	//排序完成后 遍历长度短的那个数组
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x := nums1[i]
		y := nums2[j]
		if x == y {
			//这里都用x作为基准参考
			//如果相等就说明找到 添加到数组中
			inter = append(inter, x)
			i++
			j++
		} else if x < y {
			//因为排过顺序了，但是x<y，说明x要大一点或许才能和y有相等的可能，所以将i的指针右移动

			i++
		} else {
			//反之 j 右移
			j++
		}

	}
	fmt.Println(inter)

	return
}

func main() {
	intersection()

}

//复杂度分析
/*
时间复杂度：O(mlogm + nlogn),m n 分别是数组的长度，
对两个数组进行排序的时间复杂度是O(mlogm) O(nlogn)
双指针寻找交集的时间复杂度为O(m+n),因此总时间复杂度为O(mlogm + nlogn)

*/
