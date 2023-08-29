package main

import "fmt"

//package 切片使用入门与数据共享问题处理
func main() {

	//创建切片
	//基于数组
	// 先定义一个数组
	months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	// 基于数组创建切片
	q2 := months[3:6]     // 第二季度
	summer := months[5:8] // 夏季
	all := months[:]      //基于 months 的所有元素创建切片
	all6 := months[:6]    //基于 months 的所有元素创建切片
	fmt.Println(q2)
	fmt.Println(summer)
	fmt.Println(all)
	fmt.Println(all6)
	fmt.Println(len(all6))
	fmt.Println(cap(all6))

	//基于切片
	var mySlice1 []int = make([]int, 5) //初始长度为 5 的整型切片
	fmt.Println(len(mySlice1))
	fmt.Println(cap(mySlice1))
	var mySlice2 []int = make([]int, 5, 10) // mySlice2 := make([]int, 5, 10)
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))
	//自动扩容
	mySlice2 = append(mySlice2, 8)
	fmt.Println(mySlice2)
	//内容复制
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	// 复制 slice1 到 slice 2
	copy(slice2, slice1) // 只会复制 slice1 的前3个元素到 slice2 中
	// slice2 结果: [1, 2, 3]
	// 复制 slice2 到 slice 1
	copy(slice1, slice2) // 只会复制 slice2 的 3 个元素到 slice1 的前 3 个位置
	// slice1 结果：[5, 4, 3, 4, 5]

	//动态删除元素
	// slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// slice3 = slice3[:len(slice3)-5] // 删除 slice3 尾部 5 个元素
	// slice3 = slice3[5:]             // 删除 slice3 头部 5 个元素
	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice3)
	slice4 := append(slice3[:0], slice3[3:]...) // 删除开头三个元素
	fmt.Println(slice3)
	fmt.Println(slice4)
	// slice5 := append(slice3[:1], slice3[4:]...) // 删除中间三个元素
	// slice6 := append(slice3[:0], slice3[:7]...) // 删除最后三个元素
	// slice7 := slice3[:copy(slice3, slice3[3:])] // 删除开头前三个元素
	// 创建一个切片
	mySlice := []int{1, 2, 3, 4, 5}

	// 删除索引为2的元素
	mySlice = append(mySlice[:2], mySlice[3:]...)

	fmt.Println(mySlice) // 输出：[1 2 4 5]
}
