package main

import (
    "fmt"
)

func setMap(old_m map[string]int, k string, v int) (map[string]int) {
    old_m[k] = v
    return old_m
}

func main() {
    m := make(map[string]int)
    fmt.Println(m)

    m = setMap(m, "yq", 18)
    fmt.Println(m)
    m = setMap(m, "ls", 17)
    fmt.Println(m)
    m = setMap(m, "zs", 19)
    fmt.Println(m)

    keys := []string{}
    for _k, _v := range m {
        fmt.Println(_k, _v)
        keys = append(keys, _k)
    }
    fmt.Println(keys)
    
    fmt.Println(keys)
}
