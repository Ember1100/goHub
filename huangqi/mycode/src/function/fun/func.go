package main

import "fmt"

//匿名函数的定义和使用
func main() {

	var j int = 1
	f := func() {
		var i int = 1
		fmt.Printf("s ,i, j:%q, %d, %d\n", "hello", i, j)
	}

	f()
	j += 2
	f()

	//将匿名函数作为函数参数
	add := func(a, b int) int {
		return a + b
	}
	var number int
	// 将函数类型作为参数
	func(call func(int, int) int) {
		number = call(1, 2)
	}(add)
	fmt.Println("number==============", number)
	// 普通的加法操作
	add1 := func(a, b int) int {
		return a + b
	}
	// 定义多种加法算法
	base := 10
	add2 := func(a, b int) int {
		return a*base + b
	}
	handleAdd(1, 2, add1)
	handleAdd(1, 2, add2)
}

// 将匿名函数作为参数
func handleAdd(a, b int, call func(int, int) int) {
	fmt.Println(call(a, b))
}
