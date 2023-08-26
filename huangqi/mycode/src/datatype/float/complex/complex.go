package main

import "fmt"

//package float
//复数 在 Go 语言中，复数支持两种类型：complex64（32 位实部和虚部） 和 complex128（64 位实部与虚部）
//对应的表示示例如下，和数学概念中的复数表示形式一致：
func main() {
	var complexValue1 complex64 = 1.10 + 10i // 由两个 float32 实数构成的复数类型

	// complexValue2 := 1.10 + 10i        // 和浮点型一样，默认自动推导的实数类型是 float64，所以 complexValue2 是 complex128 类型
	// complexValue3 := complex(1.10, 10) // 与 complexValue2 等价
	fmt.Println(complexValue1)

}
