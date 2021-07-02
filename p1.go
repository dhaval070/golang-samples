package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

type Person struct {
    name string
    age int
}

var sp = strings.Repeat("-", 50)

func main() {
    readStdin()
    p := &Person { "dhaval", 33 }
    fmt.Println((*p).name)

    p1 := Person { "dhaval", 33 }

    if (p1 == *p) {
        fmt.Println("structs are equal")
    }

    looptest()
    gototest()
    variadic(1, 2, 3)
    fibbo(5)

    fmt.Println(sp)

    f := plusx(5)
    fmt.Println(f(55))

    fmt.Println(sp)

    fmt.Println("Map demo")
    fn := func (n int) int {
        return n + 1
    }

    fmt.Println(Map(fn, 1, 2, 3))
}

func readStdin() {
    fmt.Print("Enter something: ")
    reader := bufio.NewReader(os.Stdin)
    line, _ := reader.ReadString('\n')

    line = strings.Replace(line, "\n", "", -1)
    fmt.Println("you entered: " + line)
    fmt.Println(sp)
}

func modify(s [3]int) {
    s[0] = 99
}

func looptest() {
    fmt.Println("looptest")
    var arr [5]int

    for i := 0; i < 5; i++ {
        arr[i] = i
    }

    fmt.Println(arr)
    fmt.Println(sp)
}

func gototest() {
    fmt.Println("gototest")
    i := 0
    L: fmt.Println(i)
        i++

        if i < 5 {
            goto L
        }
    fmt.Println(sp)
}

func variadic(a ...int) {
    fmt.Println("variadic")
    for _, n := range a {
        fmt.Println(n)
    }
    fmt.Println(sp)
}

func fibbo(n int) {
    fmt.Println("fibbo")
    a := 0
    b := 1

    for i := 0; i < n; i++ {
        c := a + b
        fmt.Println(b)

        a = b
        b = c
    }
    fmt.Println(sp)
}

func plusx(n int) func(int) int {
    fmt.Println("plusx")

    return func(x int) int {
        return n + x
    }
}

func Map(f func(int) int, a ...int) []int {
    var res = make([]int, len(a))

    for i, v := range a {
        res[i] = f(v)
    }
    return res
}
