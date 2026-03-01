package main

import (
    "errors"
    "fmt"
)

func Greetings(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }
    message := fmt.Sprintf("Hi, %v! Welcome", name)
    return message, nil
}

func main() {
    msg, err := Greetings("Alice")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(msg)
}   