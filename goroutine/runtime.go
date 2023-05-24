package main

import (
    "fmt"
    "runtime"
    "time"
)

func GOGosched() {
    go func(s string) {
        for i := 0; i < 2; i++ {
            fmt.Println(s)
        }
    }("world")
    // 主协程
    for i := 0; i < 2; i++ {
        // 切一下，再次分配任务
        runtime.Gosched()
        fmt.Println("hello")
    }
}

func GoGoexit() {
    go func() {
        defer fmt.Println("A.defer")
        func() {
            defer fmt.Println("B.defer")
            // 结束协程
            runtime.Goexit()
            defer fmt.Println("C.defer")
            fmt.Println("B")
        }()
        fmt.Println("A")
    }()
    for {
    }
}

func GoGOMAXPROCS() {
    a := func () {
        for i := 1; i < 10; i++ {
            fmt.Println("A:", i)
        }
    }

    b :=func () {
        for i := 1; i < 10; i++ {
            fmt.Println("B:", i)
        }
    }

    go a()
    go b()
    time.Sleep(time.Second)
}

func a() {
    for i := 1; i < 10; i++ {
        fmt.Println("A:", i)
    }
}

func b() {
    for i := 1; i < 10; i++ {
        fmt.Println("B:", i)
    }
}

func main() {
    // GoGoexit()
    // GOGosched()
    // runtime.GOMAXPROCS(1)
    // GoGOMAXPROCS()

    runtime.GOMAXPROCS(1)
    go a()
    go b()
    time.Sleep(time.Second)
}