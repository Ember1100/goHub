package main

import (
	"fmt"
	"strconv"
)

//package convert 基本数据类型之间的转化

func main() {
	//数值类型之间的转化 整型之间的转化
	v1 := uint(16)   // 初始化 v1 类型为 unit
	v2 := int8(v1)   // 将 v1 转化为 int8 类型并赋值给 v2
	v3 := uint16(v2) // 将 v2 转化为 uint16 类型并赋值给 v3
	fmt.Println(v3)

	//整型与浮点型之间的转化
	i1 := 99.99
	i2 := int(i1) // v2 = 99
	fmt.Println(i2)
	i3 := 99
	i4 := float64(i3) // i4 = 99
	fmt.Println(i4)

	//字符串和其他基本类型之间的转化
	v11 := 65
	v22 := string(rune(v11)) // v22 = A
	fmt.Println("v22 =", v22)
	var v33 int32 = 30028
	v44 := string(v33) // v44 = 界
	fmt.Println("v44 = ", v44)
	convert()
	//strconv 包
}

func convert() {
	v1 := "100"
	v2, _ := strconv.Atoi(v1) // 将字符串转化为整型，v2 = 100
	fmt.Println("v2 =", v2)

	v3 := 100
	v4 := strconv.Itoa(v3) // 将整型转化为字符串, v4 = "100"
	fmt.Println("v4 =", v4)

	v5 := "true"
	v6, _ := strconv.ParseBool(v5) // 将字符串转化为布尔型
	v5 = strconv.FormatBool(v6)    // 将布尔值转化为字符串
	fmt.Println("v5 =", v5)

	v7 := "100"
	v8, _ := strconv.ParseInt(v7, 10, 64) // 将字符串转化为整型，第二个参数表示进制，第三个参数表示最大位数
	v7 = strconv.FormatInt(v8, 10)        // 将整型转化为字符串，第二个参数表示进制
	fmt.Println("v7 =", v7)

	v9, _ := strconv.ParseUint(v7, 10, 64) // 将字符串转化为无符号整型，参数含义同 ParseInt
	v7 = strconv.FormatUint(v9, 10)        // 将无符号整数型转化为字符串，参数含义同 FormatInt

	v10 := "99.99"
	v11, _ := strconv.ParseFloat(v10, 64) // 将字符串转化为浮点型，第二个参数表示精度
	v10 = strconv.FormatFloat(v11, 'E', -1, 64)

	q := strconv.Quote("Hello, 世界")              // 为字符串加引号
	q = strconv.QuoteToASCII("Hello, 世界, hello") // 将字符串转化为 ASCII 编码

	fmt.Println("q =", q)
}
