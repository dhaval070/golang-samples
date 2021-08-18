package main
import "fmt"

func main() {
    i := 0
    for {
        j := i
        fmt.Println(j)
        for j > 0 {
            fmt.Printf("i: %d, j: %d\n", i, j)
            j--
        }
        i++
        if i > 10 {
            break
        }
    }

}
