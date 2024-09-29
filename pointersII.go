//Go'da bir fonksiyon işaretçisini başka bir fonksiyona argüman olarak nasıl geçireceğimiz

package main

import "fmt"

func square(x int) int {
return x * x
}

func apply(f func(int) int, x int) int {
return f(x)
}

func main() {
result := apply(square, 5)
fmt.Println("Result =", result)
}