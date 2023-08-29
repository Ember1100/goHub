package main

import (
	"fmt"
	"mycode/src/function/sort"
)

// package function
func main() {
	var array []int = []int{364, 637, 341, 406, 747, 995, 234, 971, 571, 219, 993, 407, 416, 366, 315, 301, 601, 650}
	sort.MergeSort(array)
	fmt.Println(array)
}
