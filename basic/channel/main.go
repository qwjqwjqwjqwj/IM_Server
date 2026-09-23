package main

import (
	"fmt"
	"sync"
)

func produce(out chan<- int) { //生产者 只允许往通道里写

}

func consumer(in <-chan int) { //消费者 只能读数据

}

func channel_example() {
	c := make(chan string) //无缓冲

	var wg sync.WaitGroup
	wg.Add(1) //goroutine计数器加1
	go func() {
		defer wg.Done() //匿名函数退出后 goroutine计数器减一
		defer fmt.Println("func over...")

		fmt.Println("func begin...")
		c <- "the information comes from goroutine.."

	}()
	r := <-c
	fmt.Println(r)
	wg.Wait() //等待goroutine计数器归0
}

func chan_range() {

	c := make(chan int, 5)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
	}()
	for x := range c {
		fmt.Println(x)
	}
}

func main() {

	//channel_example()
	chan_range()

}
