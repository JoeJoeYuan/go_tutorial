package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// 定义一个结构体，用于存储yaml文件中的数据
type Person struct {
	Name      string   `yaml:"name"`
	Age       int      `yaml:"age"`
	IsStudent bool     `yaml:"is_student"`
	Hobbies   []string `yaml:"hobbies"`
	Address   struct {
		Street   string `yaml:"street"`
		City     string `yaml:"city"`
		State    string `yaml:"state"`
		ZipCode  string `yaml:"zip_code"`
	} `yaml:"address"`
}

func main() {
	// 读取yaml文件
	data, err := ioutil.ReadFile("example.yaml")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	// 解析yaml文件
	var person Person
	err = yaml.Unmarshal(data, &person)
	if err != nil {
		log.Fatalf("error parsing yaml: %v", err)
	}

	// 输出解析结果
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("IsStudent: %t\n", person.IsStudent)
	fmt.Printf("Hobbies: %v\n", person.Hobbies)
	fmt.Printf("Address: %v\n", person.Address)
}
