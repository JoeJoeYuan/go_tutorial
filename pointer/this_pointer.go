package main

import (
  "fmt"
)

func modify1(x int) {
  x = 100
}

func modify2(x *int) {
  *x = 100
}

func isNil() {
  var p *string
  fmt.Println(p)
  fmt.Printf("p的值是%v\n", p)
  if p != nil {
      fmt.Println("非空")
  } else {
      fmt.Println("空值")
  }
}

// 指针1
func testNewMake() {
  var a *int
  *a = 100
  fmt.Println(*a)

  var b map[string]int
  b["测试"] = 100
  fmt.Println(b)
}

// 指针2
func testNewMake2() {
  var a int
  fmt.Println(&a)
  var p *int
  p = &a
  *p = 20
  fmt.Println(a)
}

func main() {
  a := 10
  modify1(a)
  fmt.Println(a) // 10
  modify2(&a)
  fmt.Println(a) // 100

  isNil()

  // testNewMake()

  testNewMake2()
}