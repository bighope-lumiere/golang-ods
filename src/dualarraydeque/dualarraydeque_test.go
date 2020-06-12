package dualarraydeque

import (
    "testing"
)

func TestNew(t *testing.T) {
    dad := New()
    if dad.Size() != 0 {
        t.Errorf("New DualArrayDeque has wrong size %v expected 0", dad.Size())
    }
}

func TestAdd(t *testing.T) {
    dad := New()
    dad.Add(0, 1)
    if dad.Size() != 1 {
        t.Errorf("DualArrayDeque has wrong size %v", dad.Size())
    }
    if dad.Get(0) != 1 {
        t.Errorf("DualArrayDeque[0] has wrong value %v", dad.Get(0))
    }
    dad.Add(0, -1)
    if dad.Size() != 1 {
        t.Errorf("DualArrayDeque has wrong size %v", dad.Size())
    }
    if dad.Get(0) != -1 {
        t.Errorf("DualArrayDeque[0] has wrong value %v", dad.Get(0))
    }
}

func TestGet(t *testing.T) {
    dad := New()
    for i := 0; i < 100; i++ {
        dad.Add(i, i)
    }
    for i := 0; i < 100; i++ {
        if dad.Get(i) != i {
            t.Errorf("Get(%v) returns wrong value %v, expected %v", i, dad.Get(i), i)
        }
    }
}

func TestSet(t *testing.T) {
    dad := New()
    for i := 0; i < 100; i++ {
        dad.Add(0, i)
    }
    dad.Set(88, 199)
    if dad.Get(88) != 199 {
        t.Errorf("Get(88) returns wrong value %v, expected 199", dad.Get(88))
    }
}

func TestRemove(t *testing.T) {
    dad := New()
    for i := 0; i < 100; i++ {
        dad.Add(i, i)
    }
    dad.Set(88, 199)
    dad.Add(87, 133)
    var test int
    for i := 0; i < 80; i++ {
        test = dad.Remove(0)
        if test != i {
            t.Errorf("Remove(0) returns wrong value %v, Expected %v", test, i)
        }
    }
}
