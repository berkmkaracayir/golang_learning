// adı değişen txt dosyasını sildirdik

package main

import (
    "fmt"
    "os"
)

func main() {

    // Dosya silme
    hata := os.Remove("newText.txt")
    if hata != nil {
        fmt.Println("Dosya silinirken hata oluştu:", hata)
        return
    }

    // Dosyanın var olup olmadığını kontrol etme
    if _, hata := os.Stat("newText.txt"); hata == nil {
        fmt.Println("Dosya mevcut.")
    } else {
        fmt.Println("Dosya mevcut değil.")
    }
}