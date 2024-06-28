package main

import (
	"fmt"
)

func main() {

	//无缓冲的通道
	noch := make(chan int)
	go recv(noch)
	noch <- 10

	//有缓冲的通道
	ch := make(chan int, 1)
	ch <- 10
	fmt.Println("发送成功")

	//close()
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)

	}()

	for {
		data, ok := <-c
		if ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("main结束")

	// 如何优雅的从通道循环取值

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	//  开启goroutine从ch1中接收值，并将该值的平方发送到ch2中

	go func() {
		for {
			v, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- v * v
		}
		close(ch2)
	}()

	for i := range ch2 {
		fmt.Println(i)
	}

	//单向通道
	ch11 := make(chan int)
	ch22 := make(chan int)
	go counter(ch11)
	go squarer(ch22, ch11)
	printer(ch22)
}

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
