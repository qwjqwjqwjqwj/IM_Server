package main

import (
	"fmt"
	"sync"
)

/*
// 程序启动 goroutine执行main函数 也就是main本质是goroutinue
// main的结束导致有些结果没有打印出来  ```必须同步````
func main() {
	for i := 0; i < 3; i++ {
		go fmt.Println("goroutine", i)
	}
	fmt.Println("main 结束")
}
*/

func main() {
	var wg sync.WaitGroup //计数器 每添加一个goroutine就Add(1) 结束就Done()减一 Wait()阻塞直到计数器为0
	for i := 0; i < 3; i++ {
		wg.Add(1) // 在启动 goroutine 计数器+1
		go func(n int) {
			defer wg.Done() // goroutine结束计数器   -1
			fmt.Println("goroutine", n)
		}(i) //（）是匿名函数立即执行的意思
	}

	wg.Wait() // 等所有 goroutine 完成 也就是计数器为0
	fmt.Println("main 结束")

}
