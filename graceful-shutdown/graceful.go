package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Merhaba, Go!")
}

func main() {
    srv := &http.Server{
        Addr: ":8080",
    }

    http.HandleFunc("/", handler)

    // Sinyal kanalı oluştur
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        fmt.Println("Sunucu başlatılıyor...")
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            fmt.Printf("Sunucu hatası: %v\n", err)
        }
    }()

    <-quit // Kapatma sinyali gelene kadar bekle
    fmt.Println("Sunucu kapatılıyor...")

    // Context ile sunucuyu kapatma işlemi için 5 saniye bekle
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        fmt.Printf("Sunucu zarif bir şekilde kapatılamadı: %v\n", err)
    }

    fmt.Println("Sunucu kapandı.")
}