package main

import (
	"fmt"
	"unsafe"
)

//package pointer 指针使用入门与 unsafe.Pointer

func main() {
	ac := 100
	fmt.Println(ac)
	var ptr *int = &ac // 声明指针类型,初始化指针类型值为变量 a

	fmt.Println(ptr)
	fmt.Println(*ptr)
	//通过内置函数 new 声明指针
	aa := new(int)
	*aa = 100
	fmt.Println(aa)
	//通过指针传值
	a := 1
	b := 2
	swap(&a, &b)
	fmt.Println(a, b)
	//指针类型转化
	i := 10
	var p *int = &i

	var fp *float32 = (*float32)(unsafe.Pointer(p))
	*fp = *fp * 10
	fmt.Println(i) // 100

	//指针运算实现
	arr := [4]int{1, 2, 3, 6}
	ap := &arr

	sp := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ap)) + unsafe.Sizeof(arr[0])))
	fmt.Println(*sp)
	fmt.Println(unsafe.Sizeof(arr[0]))
	*sp += 3
	fmt.Println(arr)
}

func swap(a, b *int) {
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}
