package main

import (
    "fmt"
    "os"
    "os/user"
    "mymonkey/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Hello %s! This is my version of Monkey!\n", user.Username)
    fmt.Println("Right now it only converts text into tokens, but I'm making a parser")
    repl.Start(os.Stdin, os.Stdout)
}
