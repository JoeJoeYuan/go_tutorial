package main

import (
    "flag"
    "fmt"
	"strings"
)

// 自定义flag类型
type UuFlag []string
func (u *UuFlag) String() string {
	return fmt.Sprintf("%v", *u)
}
func (u *UuFlag) Set(value string) error {
	// 逗号分割字符
	names := strings.Split(value, ",")
	*u = append(*u, names...)
	return nil
}

// flag: https://studygolang.com/static/pkgdoc/pkg/flag.htm
func main() {
    var name string
    var age int

	// 如果绑定到了某个变量，它们是值
    flag.StringVar(&name, "name", "", "Your name")
    flag.IntVar(&age, "age", 0, "Your age")

	// 如果使用的是flag本身，他们是指针
	var (
		no = flag.Int("no", 89757, "No")
		address = flag.String("address", "cd", "Your address")
		ss = flag.Bool("ss", true, "ss")
	)

	// 自定义类型
	// 需要实现
	var uu UuFlag
	flag.Var(&uu, "uu", "uu")

	// 解析
    flag.Parse()

    fmt.Printf("Name: %s\n", name)
    fmt.Printf("Age: %d\n", age)

	fmt.Printf("address: %s\n", *address)
	fmt.Printf("no: %d\n", *no)
	fmt.Printf("SS: %s\n", *ss)

	fmt.Printf("uu: %s\n", uu)

}