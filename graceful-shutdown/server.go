package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Merhaba, Go!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Sunucu başlatılıyor...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Sunucu hatası:", err)
    }
}