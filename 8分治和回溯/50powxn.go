package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1 / x
	}
	fmt.Println(x, n)
	a := myPow(x, n/2)
	b := myPow(x, n%2)
	fmt.Println("a,b result", a, b)
	fmt.Println("乘积", b*a*a)
	return b * a * a

}

func myPow2(x float64, n int) float64 {
	//正数
	if n >= 0 {
		return quickMul(x, n)
	}
	//负数
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	//终止条件
	if n == 0 {
		return 1
	}
	//下一层
	y := quickMul(x, n/2)

	//偶数
	if n%2 == 0 {
		return y * y
	}
	//如果是奇数还要多乘一个x
	return y * y * x
}

func main() {
	myPow2(2, 10)

}
