package main

import (
    "fmt"
    "os"
)

func main() {
    dosya, hata := os.Create("output.txt")
    if hata != nil {
        fmt.Println("Dosya oluşturulurken hata oluştu:", hata)
        return
    }
    defer dosya.Close()

    icerik := "Merhaba, bu dosyaya yazıldı!"
    _, hata = dosya.WriteString(icerik)
    if hata != nil {
        fmt.Println("Dosyaya yazılırken hata oluştu:", hata)
        return
    }

    fmt.Println("Veri başarıyla dosyaya yazıldı.")
}