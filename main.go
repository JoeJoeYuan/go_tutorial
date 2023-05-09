package main

import (
    "fmt"
)

// 1. array

// var arr = [4]int(1, 2, 3)
var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 2: "tom"}

func array() {
    a := [3]int{1, 2}           // 未初始化元素值为 0。
    b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
    c := [5]int{2: 100, 4: 200} // 使用引号初始化元素。
    d := [...]struct {
        name string
        age  uint8
    }{
        {"user1", 10}, // 可省略元素类型。
        {"user2", 20}, // 别忘了最后一行的逗号。
    }
    fmt.Println(arr0, arr1, arr2, str)
    fmt.Println(a, b, c, d)
}

// 求元素和
func sumArr(a [10]int) int {
    var sum int = 0
    for i := 0; i < len(a); i++ {
        sum += a[i]
    }
	fmt.Println(sum)
    return sum
}

func main() {
	array()
	a := [...]int{1,2,3,4,5,6,7,8,9,10}
	sumArr(a)
}