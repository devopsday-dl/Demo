package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello %s,this is %s,version is v1.0.1", os.Getenv("CONF"), os.Getenv("USERNAME"))
}

func main() {
    http.HandleFunc("/gotc", handler)
    log.Fatal(http.ListenAndServe(":9999", nil))
}
