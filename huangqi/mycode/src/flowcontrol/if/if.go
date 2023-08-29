package main

import "fmt"

//package if
func main() {
	score := 100
	if score > 90 {
		fmt.Println("Grade: A")
	} else if score > 80 && score <= 90 {
		fmt.Println("Grade: B")
	} else if score > 70 && score <= 80 {
		fmt.Println("Grade: C")
	} else if score > 60 && score <= 70 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}
}
