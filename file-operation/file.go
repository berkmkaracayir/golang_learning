package main

import (
    "fmt"
    "os"
)

func main() {
    dosya, hata := os.Open("example.txt")
    if hata != nil {
        fmt.Println("Dosya açılırken hata oluştu:", hata)
        return
    }
    defer dosya.Close()

    veri := make([]byte, 100)
    sayac, hata := dosya.Read(veri)
    if hata != nil {
        fmt.Println("Dosya okunurken hata oluştu:", hata)
        return
    }

    fmt.Printf("Okunan %d byte: %s\n", sayac, veri[:sayac])
}