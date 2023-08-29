package main

import "fmt"

//package array 数组使用入门及其不足
func main() {
	var a [8]byte           // 长度为8的数组，每个元素为一个字节
	var b [3][3]int         // 二维数组（9宫格）
	var c [3][3][3]float64  // 三维数组（立体的9宫格）
	var d = [3]int{1, 2, 3} // 声明时初始化
	var e = new([3]string)  // 通过 new 初始化
	fmt.Println(len(a))
	fmt.Println(len(b))
	fmt.Println(len(b))
	fmt.Println(len(c))
	fmt.Println(len(d))
	fmt.Println(len(e))
	array()
}

func array() {
	fmt.Println("****************start**********************")
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		fmt.Printf(" i=%d  vaule=%v", i, a[i])
	}
	fmt.Println()
	for i, v := range a {
		fmt.Println("Element", i, "of arr is", v)
	}

	// 通过二维数组生成九九乘法表
	var multi [9][9]string
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			n1 := i + 1
			n2 := j + 1
			if n1 < n2 { // 摒除重复的记录
				continue
			}
			multi[i][j] = fmt.Sprintf("%dx%d=%d", n2, n1, n1*n2)
		}
	}

	// 打印九九乘法表
	for _, v1 := range multi {
		for _, v2 := range v1 {
			fmt.Printf("%-8s", v2) // 位宽为8，左对齐
		}
		fmt.Println()
	}
}
