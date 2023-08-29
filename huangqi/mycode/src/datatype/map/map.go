package main

import "fmt"

//字典 map
func main() {
	//字典声明和初始化
	var testMap map[string]int = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	k := "two"
	v, ok := testMap[k]
	if ok {
		fmt.Printf("The element of key %q: %d\n", k, v)
	} else {
		fmt.Println("Not found!")
	}

	//使用入门
	//初始化切片
	var hasMap map[string]int32 = make(map[string]int32, 10)
	key := "one"
	hasMap[key] = 50
	hasMap["two"] = 100
	value, ok := hasMap[key]
	if ok {
		fmt.Printf("k = %q, value = %d\n", key, value)
	} else {
		fmt.Println("Not found!")
	}
	//删除元素
	//delete(hasMap, key)
	//遍历map
	Map := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for key, value := range Map {
		fmt.Println(key, value)
	}
}
