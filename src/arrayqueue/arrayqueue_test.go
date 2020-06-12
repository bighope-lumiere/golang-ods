package arrayqueue
import (
    "testing"
)

func TestNew(t *testing.T) {
    aq := New()
    if aq.n != 0 {
        t.Errorf("New ArrayQueue has wrong size %v", aq.n)
    }
    if aq.j != 0 {
        t.Errorf("New ArrayStack has wrong delete-index %v", aq.j)
    }
}

func TestAdd(t *testing.T) {
    aq := New()
    for i := 0; i < 10; i++ {
        aq.Add(i)
        if aq.n != i+1 {
            t.Errorf("Array Queue has wrong size %v, expected %v", aq.n, i+1)
        }
    }
}

func TestRemove(t *testing.T) {
    aq := New()
    for i := 0; i < 10; i++ {
        aq.Add(i)
    }
    var test int
    for i := 0; i < 10; i++ {
        test = aq.Remove()
        if test != i {
            t.Errorf("Remove() returns wrong value %v, Expected %v", test, i)
        }
    }
}
