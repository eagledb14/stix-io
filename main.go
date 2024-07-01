package main

import (
    // "fmt"
    "os"
)

func main() {
    // fmt.Println(read())
    read().ToYara()
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
