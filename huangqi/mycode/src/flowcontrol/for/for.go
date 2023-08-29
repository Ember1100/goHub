package main

import "fmt"

//package for
func main() {
	var arr []int32 = []int32{1, 12, 556, 45, 32}
	for i, v := range arr {
		fmt.Printf("index = %d, value = %d\n", i, v)
	}

	var ar []int32 = make([]int32, 0)
	ar = append(ar, arr...)

	for _, v := range ar {
		fmt.Printf("value = %d\n", v)
	}
	ar = append(ar, 0)
	fmt.Println(len(ar))
}
