package main

import (
    "net/http"
    "fmt"
    "os"
    "strings"
)

func main() {
    http.HandleFunc("/",
        func(w http.ResponseWriter, r *http.Request) {
            fmt.Fprintf(w, "Hello World - Path=%s", r.URL.Path[1:])
        },
    )
    port := os.Getenv("HELLO_PORT")
    if strings.TrimSpace(port) == "" {
        port = "8080"
    }
    address := fmt.Sprintf(":%s", port)
    fmt.Printf("Starting Hello-Web at %s\n", address)

    if err := http.ListenAndServe(address, nil); err != nil {
        panic(err)
    }
}

