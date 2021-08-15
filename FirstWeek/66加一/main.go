package main

var nums3 = []int{9, 9, 9, 9}
var digits = []int{5, 6, 7, 8, 9}

func plusOne() []int {
	//根据题意我们倒序循环该数组
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		//先给数组最后一个数加1
		digits[i] = digits[i] % 10
		//因为是 “加一”并且每个元素之存储单个数字也就是0～9
		//这里分三种情况：
		//1.在数字是0～8的情况下不用考虑进位问题：[2,8,4,5]
		//2.只在数组末尾一位进一问题：[1,9,4,9]
		//3.所有的数字都要进位问题：[9,9,9,9]
		if digits[i] != 0 {
			//0 < x < 10的数组，对10取余数，结果都是它本身，比如 8 % 10 = 8
			//所以对于第一种：我们在数组末尾+1后对10取余数，结果再赋值给digits[i]
			//             如果!=0，那么刚好满足题目要求的 “加一”的要求
			//第二种：在数组末尾+1后对10取余数，结果再赋值给digits[i]
			//如果=0，那么进入下一次循环，如果这时的数字<9，那么刚好是
			//该数字加1，满足了进位的要求。[1,9,4,9]=>[1,9,5,0]
			return digits
		}
	}
	//对于第三种，在上面的循环中没有一项是满足digits[i] != 0的，那么就进入下面的分支
	slice := make([]int, len(digits)+1)
	//因为每一位都要进位，那么最终的结果是要比原数组多一位的
	//初始化一个切片，默认值都是0，然后将首位赋值为1即可
	//[9,9,9,9] => [1,0,0,0,0]
	slice[0] = 1
	return slice

}

func main() {
	plusOne()
}
