package main

import (
	"sync"
)

type Mutex struct {
	name  string
	value int
	mu    sync.Mutex
}

func mutexTest(this *Mutex) { //互斥锁 对共享资源操作前 Lock(),其他线程不能读也不能写，操作完 Unlock()
	this.mu.Lock()
	this.name = "abc"
	defer this.mu.Unlock()
}

type rwMutex struct {
	name string
	rwmu sync.RWMutex
}

func RWmutexTest(this *rwMutex) {
	this.rwmu.RLock() //读锁      可并发读
	defer this.rwmu.RUnlock()

	this.rwmu.Lock() //写锁    互斥
	defer this.rwmu.Unlock()

}
func main() {

}
