package main

import (
    "fmt"
    "strconv"
    "os"
)

func main() {
    count := getCount()
    fibonacci(count)
}

func getCount() int {
    if len(os.Args) > 1 {
        arg, err := strconv.Atoi(os.Args[1])
        if err == nil {
            return arg
        }
    }
    return 16
}

func fibonacci(n int) {
    a, b := 0, 1
    for i := 0; i < n; i++ {
        fmt.Printf("%d\n", a)
        a, b = a + b, a
    }
}
