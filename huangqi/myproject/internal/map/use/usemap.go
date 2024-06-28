package main

import (
	"fmt"
	"sort"
)

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)
	fmt.Println("---------------------")
	userInfo := map[string]string{
		"username": "pprof.cn",
		"password": "123456",
	}
	fmt.Println(userInfo) //
	fmt.Println("---------------------")
	score := make(map[string]int)
	score["张三"] = 90
	score["小明"] = 100
	score["王五"] = 60
	for k, v := range score {
		fmt.Println(k, v)
	}
	delete(score, "小明") //将小明:100从map中删除
	for k, v := range score {
		fmt.Println(k, v)
	}

	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
