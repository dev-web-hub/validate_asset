package main

import (
    "encoding/json"
    "fmt"
    "os"
    "yourverb/internal"
)

func main() {
    var input internal.Input
    json.NewDecoder(os.Stdin).Decode(&input)

    if err := internal.Run(input); err != nil {
        fmt.Println("error:", err)
        os.Exit(1)
    }

    fmt.Println("ok")
}
