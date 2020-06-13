package randomqueue
import (
    "testing"
    "fmt"
)

func TestNew(t *testing.T) {
    rq := New()
    if rq.n != 0 {
        t.Errorf("New RandomQueue has wrong size %v", rq.n)
    }
    if rq.j != 0 {
        t.Errorf("New RandomStack has wrong delete-index %v", rq.j)
    }
}

func TestAdd(t *testing.T) {
    rq := New()
    for i := 0; i < 10; i++ {
        rq.Add(i)
        if rq.n != i+1 {
            t.Errorf("Random Queue has wrong size %v, expected %v", rq.n, i+1)
        }
    }
}

func TestRemove(t *testing.T) {
    rq := New()
    for i := 0; i < 10; i++ {
        rq.Add(i)
    }
    for i := 0; i < 10; i++ {
        fmt.Println(rq.Remove())
    }
}
