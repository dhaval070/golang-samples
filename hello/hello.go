package main
import (
    "fmt"
    "log"
    "example.com/greetings"
)

func main() {
    log.SetPrefix("greetings:")
    log.SetFlags(0)

    msg, err := greetings.Hello("dhaval")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(msg)

    names := []string { "dhaval", "popat", "lallu" }

    msgs, err := greetings.Hellos(names)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(msgs)
}
