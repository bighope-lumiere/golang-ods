package arraydeque
import (
    "testing"
)

func TestNew(t *testing.T) {
    ad := New()
    if ad.n != 0 {
        t.Errorf("New ArrayDeque had wrong size %v", ad.n)
    }
}

func TestAdd(t *testing.T) {
    ad := New()
    ad.Add(0, 1)
    if ad.n != 1 {
        t.Errorf("ArrayDeque had wrong size %v", ad.n)
    }
    if ad.arr[0] != 1 {
        t.Errorf("ArrayDeque[0] had wrong value %v", ad.arr[0])
    }
    ad.Add(0, -1)
    if ad.n != 1 {
        t.Errorf("ArrayDeque had wrong size %v", ad.n)
    }
    if ad.arr[0] != -1 {
        t.Errorf("ArrayDeque[0] had wrong value %v", ad.arr[0])
    }
}

func TestGet(t *testing.T) {
    ad := New()
    for i := 0; i < 100; i++ {
        ad.Add(i, i)
    }
    for i := 0; i < 100; i++ {
        if ad.Get(i) != i {
            t.Errorf("Get(%v) returns wrong value %v, Expected %v", i, ad.Get(i), i)
        }
    }
}

func TestSet(t *testing.T) {
    ad := New()
    for i := 0; i < 100; i++ {
        ad.Add(0, i)
    }
    ad.Set(88, 199)
    if ad.Get(88) != 199 {
        t.Errorf("Get(88) returns wrong value %v, Expected 199", ad.Get(88))
    }
}

func TestRemove(t *testing.T) {
    ad := New()
    for i := 0; i < 100; i++ {
        ad.Add(i, i)
    }
    ad.Set(88, 199)
    ad.Add(87, 133)
    var test int
    for i := 0; i < 80; i++ {
        test = ad.Remove(0)
        if test != i {
            t.Errorf("Remove(0) returns wrong value %v, Expected %v", test, i)
        }
    }
}
