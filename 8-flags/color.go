package main

import (
    "image/color"
    "strconv"
)

func parseColor(s string) (color.RGBA, error) {

    value, err := strconv.ParseInt(s, 16, 64)
    if err != nil {
        return color.RGBA{}, err
    }

    color := color.RGBA {
        R: uint8(value % 256),
        G: uint8(value / 256 % 256),
        B: uint8(value / 256 / 256 % 256),
    }

    return color, nil
}
