package main

import "fmt"

// 定义汽车接口
type Car interface {
    Run()
}

// 定义两种不同的汽车类型
type BenzCar struct{}
type BMWCar struct{}

// 实现汽车接口的Run方法
func (b *BenzCar) Run() {
    fmt.Println("Benz car is running")
}

func (b *BMWCar) Run() {
    fmt.Println("BMW car is running")
}

// 定义一个汽车工厂，用于创建不同类型的汽车对象
type CarFactory struct{}

// 工厂方法，根据传入的参数创建对应类型的汽车对象
func (f *CarFactory) CreateCar(carType string) Car {
    switch carType {
    case "Benz":
        return &BenzCar{}
    case "BMW":
        return &BMWCar{}
    default:
        return nil
    }
}

func main() {
    // 创建汽车工厂
    factory := &CarFactory{}

    // 使用工厂方法创建不同类型的汽车对象
    benzCar := factory.CreateCar("Benz")
    bmwCar := factory.CreateCar("BMW")

    // 调用汽车对象的Run方法
    benzCar.Run()
    bmwCar.Run()
}