package main

import "fmt"

func main() {
    var p *int
    p = new(int)
    *p = 1000
    fmt.Println(p)
    fmt.Println(*p)
}
