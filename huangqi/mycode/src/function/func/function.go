package main

import (
	"errors"
	"fmt"
	model "mycode/src/oop/struct"
	"reflect"
)

func main() {
	user := new(model.User)
	User := model.User{
		Age:  12,
		Name: "HAHA",
	}
	fmt.Println(user)
	fmt.Println(*user)
	var temp *int
	fmt.Println(temp)
	panda := new(int)
	fmt.Println(panda)
	num := make([]int, 2)
	fmt.Println(num)
	fmt.Println("****************************************")
	a, b := 4, 5
	fmt.Println(a, b)
	fmt.Println(add(a, b))

	fmt.Println(mub(&a, &b))
	fmt.Println(a, b)
	myfunc(a, b, 5, 6, 9, 23)
	myPrintf(a, b, "12", "3266", false, true, User)
	fmt.Println("****************************************")
	num3, error := sub(&a, &b)
	fmt.Println(num3, error)
	num4, err := dri(&a, &b)
	fmt.Println(num4, err)
}

// 值传参
func add(a, b int) int {
	a *= 2
	b *= 3
	return a + b
}

// 引用传参
func mub(a, b *int) int {
	*a *= 2
	*b *= 3
	return *a + *b
}

// 变长参数
func myfunc(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}

// 任意类型的变长参数（泛型）
func myPrintf(args ...interface{}) {
	for _, arg := range args {
		switch reflect.TypeOf(arg).Kind() {
		case reflect.Int:
			fmt.Println(arg, "is an int value.")
		case reflect.String:
			fmt.Printf("\"%s\" is a string value.\n", arg)
		case reflect.Array:
			fmt.Println(arg, "is an array type.")
		case reflect.Bool:
			fmt.Println(arg, "is an Bool type.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

// 多返回值
func sub(a, b *int) (int, error) {
	if *a < 0 || *b < 0 {
		err := errors.New("只支持非负整数相加")
		return 0, err
	}
	*a *= 2
	*b *= 3
	return *a - *b, nil
}

// 命名返回值
func dri(a, b *int) (c int, err error) {
	if *a < 0 || *b < 0 {
		err = errors.New("只支持非负整数相加")
		return
	}
	*a *= 2
	*b *= 3
	c = *a * *b
	return
}
