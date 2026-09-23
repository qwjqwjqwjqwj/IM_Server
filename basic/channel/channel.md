var c chan int              // 声明，值为 nil，不能用
c := make(chan int)         // 无缓冲通道
    go func() { c <- 1 }()  // 发送方阻塞，直到有人接收
    x := <-c                // 接收方阻塞，直到有人发送

c := make(chan int, 10)     // 有缓冲通道，容量 10
    c := make(chan int, 3)
    c <- 1  // 不阻塞
    c <- 2  // 不阻塞
    c <- 3  // 不阻塞
    c <- 4  // 阻塞！缓冲区满了

    <-c     // 取出一个，腾出空间
    c <- 4  // 现在不阻塞了
    
c := make(chan<- int)       // 只写通道（send-only）
c := make(<-chan int)       // 只读通道（receive-only）

var send chan<- int = make(chan int)  // 只能发送
var recv <-chan int = make(chan int)  // 只能接收

channel 阻塞与死锁