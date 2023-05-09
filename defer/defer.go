package main

import (
  "fmt"
  "net/http"
)

/*
defer特性：
    1. 关键字 defer 用于注册延迟调用。
    2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
    3. 多个defer语句，按先进后出的方式执行。
    4. defer语句中的变量，在defer声明时就决定了。
defer用途：
    1. 关闭文件句柄
    2. 锁资源释放
    3. 数据库连接释放
*/

// 延迟调用，闭包1
func add(x, y int) (z int) {
  defer func() {
      z += 100
  }()

  z = x + y
  z = z - x
  // return前才会执行defer
  return
}

func print() {
  var whatever [5]struct{}
  for i := range whatever {
      // return前才会执行defer
      // 所以输出都是4
      defer func() {
        fmt.Println(i)
      }()
  }
}

// 先进后出
func Filo () {
  var whatever [5]struct{}

  for i := range whatever {
      // defer fmt.Println(i)
      fmt.Println(i)
  }
}

// 在错误的位置使用 defer
// 当 http.Get 失败时会抛出异常。
func do() error {
  res, err := http.Get("https://baidu.com")
  // if res != nil {
  //   defer res.Body.Close()
  // }
  fmt.Println(res)
  res.Body.Close()

  if err != nil {
      return err
  }

  // ..code...

  return nil
}

func main() {
  fmt.Println(add(2, 3))
  print()
  fmt.Println("----")
  Filo()
  fmt.Println("----")
  do()
}