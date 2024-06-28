// package structureprogram

package main

import (
	"flag"
	"fmt"
	"strings"
)

const boilingF = 212.0

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g 0F or %g 0C\n", f, c)
	//输出 boiling point = 212 0F or 100 0C

	const free, boilingF2 = 32.0, 212.0

	fmt.Printf("%g F = %g C\n", free, fToC(free))           //32 F = 0 C
	fmt.Printf("%g F = %g C\n", boilingF2, fToC(boilingF2)) //212 F = 100 C
	var num = 4
	fmt.Println(num)
	var s string
	fmt.Println(s)
	var fg float64
	fmt.Println(fg)
	ss := fToC(free)

	fmt.Println(ss)

	// x := 1
	// p := &x
	// fmt.Println(*p)

	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil)
	var pp *int
	fmt.Println(pp)
	v := 1
	incr(&v)
	fmt.Println(v)
	fmt.Println(incr(&v))

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	fmt.Println(gcd(4, 3))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func incr(p *int) int {
	*p++ //递增p所指向的值； p自身保持不变
	return *p
}

func gcd(x, y int) int {
	for y != 0 {

		temp := x
		fmt.Println(y)

		x = y
		fmt.Println(x)
		y = temp % y
		// x, y = y, x%y
	}
	return x
}
func equal(a int, b int) []bool {
	// write code here
	res := make([]bool, 2)
	//地址
	if &a == &b {
		res[0] = true
	} else {
		res[0] = false
	}
	//值
	if a == b {
		res[0] = true
	} else {
		res[0] = false
	}
	return res
}

func deleteElement(s []int, index int) []int {
	// write code here
	return append(s[:index], s[index+1:]...)
}
