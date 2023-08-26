package main

import (
	"fmt"
	"math"
)

// package float
func main() {
	//float32 单精度
	var floatValue1 float32 = 1.12
	var floatValue11 float32 = 1.12
	//float32、float64参与运算 floatValue1*floatValue2
	var floatValue12 float32 = floatValue11 * floatValue1
	fmt.Println("floatValue12=", floatValue12)

	//float64 双精度
	var floatValue2 float64 = 1.12
	var floatValue22 float64 = 44.12
	var floatValue23 float64 = floatValue2 / floatValue22
	fmt.Println("floatValue23=", floatValue23)
	//float32、float64不能比较大小 floatValue23 > floatValue12

	//浮点数的精度 浮点数不是一种精确的表达方式，因为二进制无法精确表示所有十进制小数，比如 0.1、0.7
	floatValue4 := 0.1
	floatValue5 := 0.7
	floatValue6 := floatValue4 + floatValue5
	fmt.Println(floatValue6)
	//浮点数的比较
	p := 0.00001
	// 判断 floatValue1 与 floatValue2 是否相等
	if math.Dim(float64(floatValue1), floatValue2) < p {
		fmt.Println("floatValue1 和 floatValue2 相等")
	}
}
