package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64
var wg sync.WaitGroup

func add() {
	for i := 0; i < 10000; i++ {
		x = x + int64(i)
	}
	wg.Done()
}
func main() {
	wg.Add(3)
	go add()
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)

	wg.Add(3)
	go saveAdd()
	go saveAdd()
	go saveAdd()
	wg.Wait()
	fmt.Println(xx)
}

var xx int

// 互斥锁
var lock sync.Mutex

func saveAdd() {
	for i := 0; i < 10000; i++ {
		lock.Lock()
		xx = xx + i
		lock.Unlock()
	}
	wg.Done()
}

var (
	xxx int64
	//读写互斥锁
	rwlock sync.RWMutex
)

// 并发安全的map
var m = sync.Map{}

//原子操作 atomic包

// 普通版加函数
func add2() {
	// x = x + 1
	x++ // 等价于上面的操作
	wg.Done()
}

// 互斥锁版加函数
func mutexAdd() {
	lock.Lock()
	x++
	lock.Unlock()
	wg.Done()
}

// 原子操作版加函数
func atomicAdd() {
	atomic.AddInt64(&x, 1)
}
