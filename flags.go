package main
import (
    "fmt"
    "flag"
)

func main() {
    var nflag = flag.Int("n", 0, "help for n")

    flag.Parse()
    fmt.Println(*nflag)
}

