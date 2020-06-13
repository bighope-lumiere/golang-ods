package rootisharraystack

import (
    "testing"
)

func TestNew(t *testing.T) {
    ras := New()
    if ras.Size() != 0 {
        t.Errorf("New DualArrayDeque has wrong size %v expected 0", ras.Size())
    }
}

func TestAdd(t *testing.T) {
    ras := New()
    ras.Add(0, 1)
    if ras.Size() != 1 {
        t.Errorf("DualArrayDeque has wrong size %v", ras.Size())
    }
    if ras.Get(0) != 1 {
        t.Errorf("DualArrayDeque[0] has wrong value %v", ras.Get(0))
    }
    ras.Add(0, -1)
    if ras.Size() != 1 {
        t.Errorf("DualArrayDeque has wrong size %v", ras.Size())
    }
    if ras.Get(0) != -1 {
        t.Errorf("DualArrayDeque[0] has wrong value %v", ras.Get(0))
    }
}

func TestGet(t *testing.T) {
    ras := New()
    for i := 0; i < 100; i++ {
        ras.Add(i, i)
    }
    for i := 0; i < 100; i++ {
        if ras.Get(i) != i {
            t.Errorf("Get(%v) returns wrong value %v, expected %v", i, ras.Get(i), i)
        }
    }
}

func TestSet(t *testing.T) {
    ras := New()
    for i := 0; i < 100; i++ {
        ras.Add(0, i)
    }
    ras.Set(88, 199)
    if ras.Get(88) != 199 {
        t.Errorf("Get(88) returns wrong value %v, expected 199", ras.Get(88))
    }
}

func TestRemove(t *testing.T) {
    ras := New()
    for i := 0; i < 100; i++ {
        ras.Add(i, i)
    }
    ras.Set(88, 199)
    ras.Add(87, 133)
    var test int
    for i := 0; i < 80; i++ {
        test = ras.Remove(0)
        if test != i {
            t.Errorf("Remove(0) returns wrong value %v, Expected %v", test, i)
        }
    }
}
