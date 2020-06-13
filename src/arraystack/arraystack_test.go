package arraystack

import (
    "testing"
)

func TestNew(t *testing.T) {
    as := New()
    if as.n != 0 {
        t.Errorf("New ArrayStack has wrong size %v", as.n)
    }
}

func TestAdd(t *testing.T) {
    as := New()
    as.Add(0, 1)
    if as.n != 1 {
        t.Errorf("ArrayStack has wrong size %v", as.n)
    }
    if as.arr[0] != 1 {
        t.Errorf("ArrayStack[0] has wrong value %v", as.arr[0])
    }
    as.Add(0, -1)
    if as.n != 1 {
        t.Errorf("ArrayStack has wrong size %v", as.n)
    }
    if as.arr[0] != -1 {
        t.Errorf("ArrayStack[0] has wrong value %v", as.arr[0])
    }
}

func TestGet(t *testing.T) {
    as := New()
    for i := 0; i < 100; i++ {
        as.Add(i, i)
    }
    for i := 0; i < 100; i++ {
        if as.Get(i) != i {
            t.Errorf("Get(%v) returns wrong value %v, Expected %v", i, as.Get(i), i)
        }
    }
}

func TestSet(t *testing.T) {
    as := New()
    for i := 0; i < 100; i++ {
        as.Add(0, i)
    }
    as.Set(88, 199)
    if as.Get(88) != 199 {
        t.Errorf("Get(88) returns wrong value %v, Expected 199", as.Get(88))
    }
}

func TestRemove(t *testing.T) {
    as := New()
    for i := 0; i < 100; i++ {
        as.Add(i, i)
    }
    as.Set(88, 199)
    as.Add(87, 133)
    var test int
    for i := 0; i < 80; i++ {
        test = as.Remove(0)
        if test != i {
            t.Errorf("Remove(0) returns wrong value %v, Expected %v", test, i)
        }
    }
}

func TestAddAll(t *testing.T) {
    as := New()
    c1 := New()
    for i := 0; i < 10; i++ {
        c1.Add(i, i+1)
    }
    c2 := New()
    for i := 0; i < 5; i++ {
        c2.Add(i, 2*(i+1))
    }
    as.AddAll(0, c1)
    as.AddAll(8, c2)
    if as.Get(7) != 8 {
        t.Errorf("Get(7) returns wrong value %v, Expected 8", as.Get(7))
    }
    if as.Get(8) != 2 {
        t.Errorf("Get(8) returns wrong value %v, Expected 2", as.Get(8))
    }
}
