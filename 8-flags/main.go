package main

import (
    "fmt"
    "log"
)

func main() {
    config, err := parseArgs()
    if err != nil {
        log.Fatal(err)
    }

    draw(config)
}

func draw(cfg Config) {
    fmt.Printf("Drawing with %v\n", cfg.String())
}
