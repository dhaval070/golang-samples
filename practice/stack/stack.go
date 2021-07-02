package stack
import "fmt"

func init() {
    fmt.Println("stack init")
}

type Stack struct {
    i int
    data [10] int
}

func (s *Stack) Push(n int) {
    s.data[s.i] = n
    s.i++
}

func (s *Stack) Pop() int {
    s.i--
    ret := s.data[s.i]
    s.data[s.i] = 0

    return ret
}

