package main

import (
	"fmt"
	"time"
)

func timeAfter() {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "任务完成"
	}()
	select {
	case res := <-ch:
		fmt.Println("收到结果：", res)
	case <-time.After(2 * time.Second): // 2秒超时 Go1.23GC会回收没触发的timer但是高频场景尽量手动服用Timer
		fmt.Println("操作超时！")
	}
}

func timer_EX() {
	timer := time.NewTimer(1 * time.Second) //定义后开始计时
	defer timer.Stop()

	ch := make(chan int)

	// 另一个goroutine，持续每隔500ms发消息（模拟持续消息流）
	go func() {
		i := 0
		for {
			time.Sleep(2000 * time.Millisecond)
			ch <- i
			i++

		}
	}()
	for {
		// ✅ 安全重置定时器
		if !timer.Reset(1 * time.Second) { //定时未到1S 返回true 跳过{} 并重置定时器
			<-timer.C //定时器到1S 返回false 进入{}清空timer.C并重新计时
		}
		select {
		case msg := <-ch:
			fmt.Println("Mission Success!", msg)
		case <-timer.C:
			fmt.Println("空闲超时")
		}
	}
}

func timerC_EX() {
	timer := time.NewTimer(1 * time.Second)
	tVal := <-timer.C
	fmt.Println(tVal)
	defer timer.Stop() //如果定时器未到时间但是调用 timer.Stop () Timer.C是空的 读取Timer.C会死锁
}
func main() {
	now := time.Now()
	fmt.Println(now)
	time.Sleep(1 * time.Second)

	timer_EX()

}
