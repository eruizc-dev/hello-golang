package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    argc := len(os.Args)
    if argc != 2 {
        fmt.Printf("Unexpected amount of arguments %d\n", argc)
        return
    }
    file := os.Args[1]
    data, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Printf("Error %s\n", err.Error())
        return
    }
    fmt.Print(string(data))
}

