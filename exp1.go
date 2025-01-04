package main

import "fmt"

func main() {
    var x int = 42
    fmt.Println("x değişkeninin değeri:", x)
    fmt.Println("x değişkeninin bellek adresi:", &x)

    var ptr *int = new(int)
    *ptr = 100
    fmt.Println("ptr işaretçisinin değeri:", *ptr)
    fmt.Println("ptr işaretçisinin bellek adresi:", ptr)
}