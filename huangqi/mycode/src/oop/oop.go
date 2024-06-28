package main

import (
	"fmt"
	inter "mycode/src/oop/interface"
	model "mycode/src/oop/struct"
)

func main() {
	student := model.Student{
		Name:  "Ember",
		Age:   23,
		Grade: 18,
	}
	var people inter.People = student
	isRun := people.Run()
	fmt.Println(isRun)
	stu := model.NewStudent("学院君", 100, 1)
	fmt.Println(*stu)
}
