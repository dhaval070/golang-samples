package stack
import "testing"
import "os"

func TestMain(m *testing.M) {
    code := m.Run()
    os.Exit(code)
}

func TestPush(t *testing.T) {
    t.Log("test push")
    var s Stack
    s.Push(2)

    if (s.data[0] != 2) {
        t.Fatal("stack head no 2")
    }
}

func TestPop(t *testing.T) {
    t.Log("test pop")
    var s Stack

    table := []int { 1, 2, 3, 4, 5, 6 }

    for _, v := range table {
        s.Push(v)
    }

    for j := len(table) - 1; j >= 0; j-- {
        n := s.Pop()

        if (table[j] != n ) {
            t.Error("it failed")
        }
    }
}
