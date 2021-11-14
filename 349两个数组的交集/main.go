package main

import "fmt"

var nums = []int{3, 4, 5, 6}
var nums2 = []int{7, 6, 5, 4, 3, 2}

func intersection() (inter []int) {
	//将两个数组分别遍历，存放进map中
	m1 := map[int]struct{}{}
	for i := 0; i < len(nums); i++ {
		m1[nums[i]] = struct{}{}
	}
	fmt.Println("this is m1 ", m1)

	m2 := map[int]struct{}{}
	for i := 0; i < len(nums2); i++ {
		m2[nums2[i]] = struct{}{}
	}
	fmt.Println(m2)
	//遍历长度短的那一个map，看长度长的map中有没有和他一样的元素，遍历短的map时间复杂度低
	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}
	//迭代map
	for v := range m1 {
		fmt.Println(v)
		if _, ok := m2[v]; ok {
			inter = append(inter, v)

		}
	}
	return

}

func main() {
	intersection()

}

//复杂度分析：
/*
时间复杂度：O(m+n)。其中m n分别是两个数组的长度。
使用两个map来分别存储两个数组中的元素需要O(m+n)的时间，
遍历较小的集合判断元素是否存在于另一个map需要O(min(m,n))的时间，
因此总的时间复杂度是O(m+n)
*/
