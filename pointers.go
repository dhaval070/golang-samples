package main

import "fmt"

func main() {
    var n = 10
    var pn *int = &n
    var zn *int

    fmt.Println(pn)
    fmt.Println(*pn)
    fmt.Println(zn)
    // following will generate seg fault
    // fmt.Println(*zn)

    size := new(int)
    fmt.Printf("size: %v, value %v, type %T\n", size, *size, size)
}
