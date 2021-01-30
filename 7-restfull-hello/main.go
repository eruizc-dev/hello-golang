package main

import (
    "fmt"
    "log"
    "net/http"
)
 
func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Println("Got a new request")
    fmt.Fprintf(w, "Hello from the go backend!")
}
 
func handleRequests() {
    http.HandleFunc("/", homePage)
    err := http.ListenAndServe(":8000", nil)
    if err != nil {
        log.Fatal(err)
    }
}
 
func main() {
    handleRequests()
}

