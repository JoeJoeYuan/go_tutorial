package main

import (
    "fmt"
)

type Hello struct {
    desc string
}

func (h *Hello) say(name string) {
    fmt.Println(h.desc, name)
}

func main() {
    // 结构体方法
    // var h Hello
    h := &Hello{
        desc: "hello",
    }
    h.say("yq")
    h = new(Hello)
    h.desc = "nihao"
    h.say("qf")
}
