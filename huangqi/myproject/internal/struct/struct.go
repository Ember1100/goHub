package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

type NewInt int // 类型定义

type MyInt = int // 类型别名

type people person

var peo person

func main() {
	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int

	var p people
	fmt.Println(p)
	fmt.Println(peo)
	var p2 = new(person)
	p2.name = "测试"
	p2.age = 18
	p2.city = "北京"
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"测试", city:"北京", age:18}
}
