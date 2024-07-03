package main

import (
    "fmt"
    "os"
)

func main() {
    // fmt.Println(read())
    fmt.Println("hi")
    // fmt.Println(read().ToYara().File())
    fmt.Println(read().ToYara().File())
    // for _, i := range read().object {
    //     fmt.println(i.pattern)
    // }
}

func read() Bundle {
    content, err := os.ReadFile("input.txt")
    if err != nil {
        panic("unkown file name")
    }
    return Unmarshall(string(content))
}
