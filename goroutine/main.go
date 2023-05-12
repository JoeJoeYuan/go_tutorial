package main

import (
    "fmt"
    "time"
    "sync"
)

func hello() {
    fmt.Println("Hello Goroutine!")
}

// 如果主协程退出了，其他任务也会退出
func mainExit() {
    // 合起来写
    go func() {
        i := 0
        for {
            i++
            fmt.Printf("new goroutine: i = %d\n", i)
            time.Sleep(time.Second)
        }
    }()
    i := 0
    for {
        i++
        fmt.Printf("main goroutine: i = %d\n", i)
        time.Sleep(time.Second)
        if i == 2 {
            // 主协程退出
            break
        }
    }
}

var wg sync.WaitGroup
func helloDefer(i int) {
    defer wg.Done() // goroutine结束就登记-1
    fmt.Println("Hello Goroutine!", i)
}

func multiGos() {
    for i := 0; i < 10; i++ {
        wg.Add(1) // 启动一个goroutine就登记+1
        go helloDefer(i)
    }
    wg.Wait() // 等待所有登记的goroutine都结束
}

func main() {
    hello()
    fmt.Println("main goroutine done!")
    fmt.Println("-----")
    // 新启动一个goroutine
    // 并发
    go hello()
    fmt.Println("main goroutine done!")
    time.Sleep(time.Second)
    fmt.Println("-----")

    // 如果主协程退出了，其他任务还执行吗
    mainExit()
    fmt.Println("-----")

    // 启动多个goroutine
    multiGos()
    fmt.Println("-----")
}
