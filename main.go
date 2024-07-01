package main

import (
    "fmt"
    "os"
)

func main() {
    // fmt.Println(read())
    for _, i := range read().Object {
        fmt.Println(i.Pattern)
    }
}

func read() Bundle {
    content, err := os.ReadFile("input.txt")
    if err != nil {
        panic("unkown file name")
    }
    return Unmarshall(string(content))
}
