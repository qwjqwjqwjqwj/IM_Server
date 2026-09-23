````Do not communicate by sharing memory; instead, share memory by communicating.````

多个线程访问一个内存 ❌️
    var counter int
    var mu sync.Mutex

    func worker() {
        mu.Lock()
        counter++      // 通过读写共享变量 counter 来“通信”
        mu.Unlock()
    } 

数据在 goroutine 之间传递 ✅️

ch := make(chan int)

// 一个专门的 goroutine 拥有这个 counter
go func() {
    counter := 0
    for delta := range ch {
        counter += delta
    }
}()

ch <- 1   // 通过 channel 把“加 1”这个意图传过去
ch <- 1