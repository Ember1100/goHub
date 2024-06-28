package main

import "fmt"

func sort(nums []int) {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

}

func Test() {
	num := []int{5, 1, 9, 3, 6}
	sort(num)
	for _, v := range num {
		fmt.Println(v)
	}
}

func main() {
	Test()
}
