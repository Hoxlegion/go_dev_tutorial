package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    names := []string{"Hank", "Berta", "Bob"}
 //  message, err := greetings.Hello("Billy")
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }

    for name, message := range messages {
        fmt.Println(name, message)
    }
    fmt.Println(messages)
}
