package main

import "fmt"

//package integer 整形默认值为0

func main() {
	//int8 带符号8位 -128~127
	var number1 int8 = -128
	fmt.Println("int8类型，最小值 =", number1)
	number1 = 127
	fmt.Println("int8类型，最大值 =", number1)

	//uint8 无符号8位 0~255
	var number2 uint8 = 0
	fmt.Println("uint8类型，最小值 =", number2)
	number2 = 255
	fmt.Println("uint8类型，最大值 =", number2)

	//int16 带符号16位 -32768~32767
	var number3 int16 = -32768
	fmt.Println("int16类型，最小值 =", number3)
	number3 = 32767
	fmt.Println("int16类型，最大值 =", number3)

	//uint16 无符号16位 0~65535
	var number4 uint16 = 0
	fmt.Println("uint16类型，最小值 =", number4)
	number4 = 65535
	fmt.Println("uint16类型，最大值 =", number4)

	//int32 带符号32位 与 rune 类型等价 -2147483648~2147483647
	var number5 int32 = -2147483648
	fmt.Println("int32类型，最小值 =", number5)
	number5 = 2147483647
	fmt.Println("int32类型，最大值 =", number5)

	//uint32 无符号32位  0~4294967295
	var number6 uint32 = 0
	fmt.Println("uint32类型，最小值 =", number6)
	number6 = 4294967295
	fmt.Println("uint32类型，最大值 =", number6)

	//int64 带符号64位 -9223372036854775808~9223372036854775807
	var number7 int64 = -9223372036854775808
	fmt.Println("int64类型，最小值 =", number7)
	number7 = 9223372036854775807
	fmt.Println("int64类型，最大值 =", number7)

	//uint64 无符号64位 0~18446744073709551615
	var number8 uint64 = 0
	fmt.Println("uint64类型，最小值 =", number8)
	number8 = 18446744073709551615
	fmt.Println("uint64类型，最大值 =", number8)

	//int 32位或64位 与具体平台相关
	var number9 int = -9223372036854775808
	fmt.Println("int类型，最小值 =", number9)
	number9 = 9223372036854775807
	fmt.Println("int类型，最大值 =", number9)

	//uint 32位或64位 与具体平台相关
	var number10 uint = 0
	fmt.Println("uint类型，最小值 =", number10)
	number10 = 18446744073709551615
	fmt.Println("uint类型，最大值 =", number10)

	//uintptr 无符号整型 32位平台下为4字节，64位平台下为8字节
	var number11 uintptr = -0
	fmt.Println("uintptr类型，最小值 =", number11)
	number11 = 18446744073709551615
	fmt.Println("uintptr类型，最大值 =", number11)
}
