package main

import "fmt"

//package string

func main() {
	//声明和初始化
	var str string = "hello word"
	fmt.Println(str)
	st := "Hello Word"
	fmt.Println(st)
	//格式化输出
	fmt.Printf("The length of \"%s\" is %d \n", str, len(str))
	results := `Search results for "Golang":
	- Go
	- Golang
	Golang Programming
	`
	fmt.Printf("%s", results)
	//将 Unicode 编码转化为可打印字符
	str1 := "Hello, 世界"
	for i, ch := range str1 {
		fmt.Println(i, string(ch))
	}
	//字符串拼接
	str = str + ", 学院君"
	fmt.Println(str)
	//字符串切片
	//Go 切片区间可以对比数学中的区间概念来理
	str2 := str[:5]  // 获取索引5（不含）之前的子串
	str3 := str[7:]  // 获取索引7（含）之后的子串
	str4 := str[0:5] // 获取从索引0（含）到索引5（不含）之间的子串
	fmt.Println("str1:", str2)
	fmt.Println("str2:", str3)
	fmt.Println("str3:", str4)

}
