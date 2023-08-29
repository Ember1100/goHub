package main

import "fmt"

// package break

func main() {
	var arr []int32 = []int32{1, 12, 556, 45, 32}
	for i, v := range arr {
		fmt.Printf("index = %d, value = %d\n", i, v)
		if arr[i] == 556 {
			fmt.Printf("index = %d, value = %d\n", i, v)
			break
		}
	}
	var arr2 []int32 = []int32{1, 12, 556, 45, 32}
	for i, v := range arr2 {
		if arr2[i] == 556 {
			continue
		}
		fmt.Printf("index = %d, value = %d\n", i, v)
	}
	arr3 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := arr3[i][j]
			if j > 1 {
				goto EXIT
			}
			fmt.Println(num)
		}
	}

EXIT:
	fmt.Println("Exit.")
}
