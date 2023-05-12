package main

import "fmt"

type BaseBird struct {
    age int
}

func (this *BaseBird) Cal()  {
    this.Add()
}

func (this *BaseBird) Add()  {
    fmt.Printf("before add: age=%d\n", this.age)
    this.age = this.age + 1
    fmt.Printf("after add: age=%d\n", this.age)
}

type DerivedBird struct {
    BaseBird
}
func (this *DerivedBird) Add()  {
    fmt.Printf("before add: age=%d\n", this.age)
    this.age = this.age + 2
    fmt.Printf("after add: age=%d\n", this.age)
}

func main() {
    // var b1 BaseBird
    // var b2 DerivedBird

    // b1 = BaseBird{age: 1}
    // b1.Cal()

    // b2 = DerivedBird{BaseBird{1}}
    // b2.Cal()


    // 2.
    p := person{"zhangsan", 22}
    s := student{person{"lisi", 20}, "000"}
    t := teacher{person{"wangwu", 25}, "000"}

    fmt.Println(s.name)			//访问【匿名】结构体的字段
    fmt.Println(t.p.name)		//访问【有名】结构体的字段。不是继承，需要指定结构体
    p.run()
    s.run()

}

type person struct {
	name string
	age int

}

//如果一个struct嵌套了另一个匿名结构体，就可以直接访问匿名结构体的字段或方法，从而实现继承
type student struct {
	person	//匿名字段，struct
	mobile string
}

//如果一个struct嵌套了另一个【有名】的结构体，叫做组合
type teacher struct {
	p person	//有名字段，struct
	mobile string
}

func (p *person) run(){
	fmt.Println(p.name, " run")
}

func (p *person) reading(){
	fmt.Println(p.name, " reading")
}

func (s *student) reading(){
	fmt.Println(s.name, " reading")
}
