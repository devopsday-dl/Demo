package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func get_hostname() string {
    hostname, _ := os.Hostname()
    return hostname
}

func get_env_var() string {
    version := os.Getenv("VERSION")
    return version
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello DevOps, pod name is %s, version is %s",get_hostname(), get_env_var())
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":9999", nil))
}
