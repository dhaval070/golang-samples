package main

import (
    "fmt"
    "math/rand"
    "sync"
    "log"
//   "runtime"
//   "time"
)
func main() {
    dmd := make([]int, 2, 5)
    log.Println(len(dmd))
    dmd = append(dmd, 11, 12, 13,14,15,16)
    fmt.Println(dmd)

    m1 := make(map[string][]int )
    m1["xx"] = []int{ 1, 2 }

    fmt.Println(m1)

//----------------------------------
    a1 := []int { 1, 2 }
    a1 = append(a1, 3, 4, 5 )

    a2 := []int { 10, 11, 12 }
    a1 = append(a1, a2...)

    fmt.Println(a1)

//----------------------------------
    x, y := 2, 3
    fmt.Println(power(x, y))

//----------------------------------
    var a = [...]int{ 1, 2, 3, 4, 5, 6, 7 }
    fmt.Println(a)

    sqr(a[2:3])
    fmt.Println(a)

//----------------------------------
    s := "dhaval"
    sr := []rune(s)
    sr[0] = 'x'
//----------------------------------
    var m [10][10]int

    for i:=0; i < 10; i++ {
        for j:=0; j < 10; j++ {
            m[i][j] = rand.Intn(10)
        }
    }

    printMap(m[:])

    var wg sync.WaitGroup
    wg.Add(len(m))

    for i := 0; i < len(m); i++ {
        var n = i
        go func() {
            defer wg.Done()
            sqr(m[n][:])
        }()
    }
    wg.Wait()
    printMap(m[:])
    fmt.Println("bye")
}
// named return type demo
func power(n, p int) (result int) {
    for result = n; p > 1; p-- {
        result *= n
    }
    return
}

func sqr(p []int) {
    for i, v := range p {
        p[i] = v * v
    }
}


func printMap(m [][10]int) {
    fmt.Println("-----------")
    for i:=0; i < len(m); i++ {

        for j:=0; j < len(m[i]); j++ {
            fmt.Printf("%3d ", m[i][j])
        }
        fmt.Print("\n")
    }
}
