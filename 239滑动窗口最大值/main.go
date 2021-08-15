package main

import "fmt"

var a = []int{1, 3, -1, -3, 5, 3, 6, 7}

func maxSlidingWindow1(nums []int, k int) []int {
	var window []int
	l, r := 0, 0
	var result []int
	n := len(nums)
	for ; r < n; r++ { //滑动窗口右边界 < 数组的长度
		if len(window) == 0 {
			window = append(window, r) //这里的r是索引数；如果window的长度=0，那么就开始放入第一个元素
		} else if nums[r] > nums[window[len(window)-1]] { //因为window是单调递减队列，因此，在window不为空的情况下要将进入window的元素和window队尾的元素进行比较
			//如果将要进来的元素 < window队尾元素，那么就正常放进来
			//如果将要进来的元素 > window队尾元素，那么需要将尾巴的元素移出去，将这个大的元素放进来
			for nums[r] > nums[window[len(window)-1]] { //将上一个小的队尾巴元素移除后，还需要和window中之前的元素比大小，如果大了，还需要移除window里面的元素，反复如此
				window = window[:len(window)-1] //左闭右开，刚好去掉window最后一个元素
				if len(window) == 0 {           //移除元素之后有可能window 变成空的，那么就需要跳出这层循环重新赋值r
					break
				}
			}
			window = append(window, r)
		} else {
			window = append(window, r) //如果window的尾巴元素大于将要进入的元素，那么就正常进入
		}
		if r > k-1 { //因为题目window的大小是3，因此当window右边界减去k要大于1，那么为了位置window是3，那么左边界就需要向右边移动
			l++
		}
		if window[0] < l { //当队首元素的下标小于滑动window，表示队首元素已经不再滑动窗口内，因此将其从window队首移除
			window = window[1:]
		}
		if r >= k-1 { //因为是单调递减队列，又因为window的大小始终为3，所以 window的index=0 就是当前滑动窗口最大的值
			result = append(result, nums[window[0]])
		}
		if window[0]<l{
			window = window[1:]
		}
		if r>k-1{
			result = append(result,r)
		}

	}
	fmt.Println(result)

	return result
}

func main() {
	maxSlidingWindow1(a, 3)

}
