package main

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	left := 2        //设置作边界为2
	right := num / 2 //设置右边界为num/2
	for left <= right {
		x := left + (right-left)/2
		tmp := x * x
		if tmp == num {
			return true
		}
		if tmp > num {
			right = x - 1
		} else {
			left = x + 1
		}

	}
	return false
}

func main() {
	isPerfectSquare(8)

}
