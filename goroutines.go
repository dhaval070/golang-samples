/*
This program creates an array of random integers and retrieves following using goroutines:
1) sum of values
2) min and max values
*/
package main
import (
    "fmt"
    "math/rand"
    "time"
)

func doSum(a []int, done chan int) {
    var s = 0

    for i:=0; i < len(a); i++ {
        fmt.Println(a[i])
        s += a[i]
    }
    done <- s
}

type MinMax struct {
    Min int
    Max int
}

func getMinMax(a []int, done chan MinMax) {
    result := MinMax { 501, -1 }

    for _, v := range a {
        if v < result.Min {
            result.Min = v
        }

        if v > result.Max {
            result.Max = v
        }
    }

    done <- result
}

func main() {
    var arr []int
    var sum int
    var minmax MinMax

    rand.Seed(time.Now().UnixNano())
    done := make(chan int)
    done1 := make(chan MinMax)

    for n:=1; n < 500; n++ {
        arr = append(arr, rand.Intn(400) + 100)
    }

    go doSum(arr, done)
    go getMinMax(arr, done1)

    sum = <- done
    minmax = <- done1

    fmt.Println(sum)
    fmt.Println(minmax)
}
