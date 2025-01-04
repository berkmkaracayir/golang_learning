package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            http.ServeFile(w, r, "public/index.html")
            return
        }
        http.ServeFile(w, r, "public"+r.URL.Path)
    })

    fmt.Println("Starting server on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}