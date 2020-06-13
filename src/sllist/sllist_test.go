package sllist
import (
    "testing"
)

func TestNew(t *testing.T) {
    l := New()
    if l.n != 0 {
        t.Errorf("New SLList has wrong size %v", l.n)
    }
}

func TestPushPop(t *testing.T) {
    l := New()
    for i := 0; i < 10; i++ {
        l.Push(i)
    }
    for i := 0; i < 10; i++ {
        test := l.Pop()
        if test != 9-i {
            t.Errorf("Pop() returns wrong value %v, expecting %v", test, 9-i)
        }
    }
}

func TestAdd(t *testing.T) {
    l := New()
    for i := 0; i < 10; i++ {
        l.Add(i)
    }
    for i := 0; i < 10; i++ {
        test := l.Remove()
        if test != i {
            t.Errorf("Pop() returns wrong value %v, expecting %v", test, i)
        }
    }
}
