package main

import (
    "fmt"
)

func main() {
    var a, b, result float64
    var operator string

    fmt.Print(">>> ")
    fmt.Scanf("%f %s %f", &a, &operator, &b)

    switch operator {
        case "+":
            result = a + b
            break
        case "-":
            result = a - b
            break
        case "*":
            result = a * b
            break
        case "/":
            result = a / b
            break
        default:
            fmt.Printf("Unknown operator %s\n", operator)
    }

    fmt.Printf("%f\n", result)
}
