// txt dosyasının adını değiştirir
package main

import (
    "fmt"
    "os"
)

func main() {
    // Dosya adını değiştirme
    hata := os.Rename("oldText.txt", "newText.txt")
    if hata != nil {
        fmt.Println("Dosya adı değiştirilirken hata oluştu:", hata)
        return
    }
}