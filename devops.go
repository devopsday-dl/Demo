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

// func get_env_var() string {
//     version := os.Getenv("VERSION")
//     return version
// }

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello GitOps,this is xiaomage,version is v1, pod name is %s",get_hostname())
}

func main() {
    http.HandleFunc("/devopsday", handler)
    log.Fatal(http.ListenAndServe(":9999", nil))
}
