package main

import "fmt"

// 定义一个接口类型
type Animal interface {
    speak() string
}

// 定义一个结构体类型
type Dog struct{}

// 实现Animal接口的speak方法
func (d Dog) speak() string {
    return "Woof!"
}

// 定义另一个结构体类型
type Cat struct{}

// 实现Animal接口的speak方法
func (c Cat) speak() string {
    return "Meow!"
}

func main() {
    // 创建一个Dog对象
    d := Dog{}

    // 创建一个Cat对象
    c := Cat{}

    // 创建一个Animal类型的切片
    animals := []Animal{d, c}

    // 遍历切片并调用每个对象的speak方法
    for _, animal := range animals {
        fmt.Println(animal.speak())
    }
}