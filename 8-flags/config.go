package main

import (
    "image/color"
    "flag"
    "fmt"
)

type Config struct {
    backgroundColor color.Color
    foregroundColor color.Color
}

func parseArgs() (Config, error) {
    background := flag.String("bg", "000000", "background color")
    foreground := flag.String("fg", "ffffff", "foreground color")
    flag.Parse()

    bgColor, err := parseColor(*background)
    if err != nil {
        return Config{}, err
    }
    fgColor, err := parseColor(*foreground)
    if err != nil {
        return Config{}, err
    }

    config := Config {
        backgroundColor: bgColor,
        foregroundColor: fgColor,
    }

    return config, nil
}

func (c *Config) String() string {
    if c == nil {
        return "nil"
    }

    return fmt.Sprintf(
        "background: %v, foreground: %v",
        c.backgroundColor,
        c.foregroundColor,
    )
}
