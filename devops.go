package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello GitOps,this is xiaomage,version is v1.0.0")
}

func main() {
    http.HandleFunc("/gotc", handler)
    log.Fatal(http.ListenAndServe(":9999", nil))
}
