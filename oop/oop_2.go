package main

import (
  "fmt"
)

type Hero struct {
  Name string
  Age uint64
}

func NewHero() *Hero {
  return &Hero{
      Name: "盖伦",
      Age: 18,
  }
}

func (h Hero) GetName() string {
  return h.Name
}

func (h Hero) GetAge() uint64 {
  return h.Age
}


func main()  {
  // t := Hero{"yq", 17}
  t := 10
  fmt.Println(t)
  t_this := &t
  fmt.Println(&t, &t_this)

  fmt.Println(t, *t_this)

  // fmt.Println(t.GetName())
  // fmt.Println(t.GetAge())
  h := NewHero()
  fmt.Println(h.GetName())
  fmt.Println(h.GetAge())
}