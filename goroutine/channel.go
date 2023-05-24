package main

import (
    "fmt"
    // "runtime"
    // "time"
)

func GoSelect() {
    var c1, c2, c3 chan int
    var i1, i2 int
    select {
        case i1 = <-c1:
            fmt.Printf("received ", i1, " from c1\n")
        case c2 <- i2:
            fmt.Printf("sent ", i2, " to c2\n")
        case i3, ok := (<-c3):  // same as: i3, ok := <-c3
            if ok {
                fmt.Printf("received ", i3, " from c3\n")
            } else {
                fmt.Printf("c3 is closed\n")
            }
        default:
            fmt.Printf("no communication\n")
    }
}

func recv(c chan int) {
    ret := <-c
    fmt.Println("接收成功", ret)
}
func main() {
    // 创建无缓冲的channel
    ch := make(chan int)

    // 无缓冲的通道必须有接收才能发送
    // 因此，无缓冲通道也被称为同步通道
    go recv(ch) // 启用goroutine从通道接收值
    ch <- 10
    fmt.Println("发送成功")
    fmt.Println("----")

    // 创建一个容量为1的有缓冲区通道
    ch_cache := make(chan int, 2)

    // 第1次发送
    ch_cache <- 20
    fmt.Println("发送成功")

    // 第2次发送
    ch_cache <- 30
    fmt.Println("发送成功")

    // 第1次接收: 20
    recv(ch_cache)
    // 第2次接收: 30
    recv(ch_cache)

    GoSelect()
}
