package main

import "fmt"

func carp(a, b int) int {
    return a * b
}

func islemUygula(islem func(int, int) int, a, b int) int {
    return islem(a, b)
}

func main() {
    sonuc := islemUygula(carp, 5, 3)
    fmt.Println(sonuc) // Çıktı: 15
}