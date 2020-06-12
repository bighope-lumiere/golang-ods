package dualarraydeque

import (
    "golang-ods/src/arraystack"
)

type DualArrayDeque struct {
    front arraystack.ArrayStack
    back arraystack.ArrayStack
}

func New() DualArrayDeque {
    return DualArrayDeque{
        front: arraystack.New()
        back: arraystack.New()
    }
}

func (this DualArrayDeque) Size() int {
    return this.front.Size() + this.back.Size()
}

func (this DualArrayDeque) Get(i int) int {
    if i < this.front.Size() {
        return this.front.Get(this.front.Size() - i - 1)
    } else {
        return this.back.Get(i - this.front.Size())
    }
}

func (this *DualArrayDeque) Set(i, x int) int {
    if i < this.front.Size() {
        return this.front.Set(this.front.Size() - i - 1, x)
    } else {
        return this.back.Set(i - this.front.Size(), x)
    }
}

func (this *DualArrayDeque) Add(i, x int) {
    if i < this.front.Size() {
        this.front.Add(this.front.Size() - i, x)
    } else {
        this.back.Add(i - this.front.Size(), x)
    }
    this.Balance()
}

func (this *DualArrayDeque) Remove(i int) int {
    var x int
    if i < this.front.Size() {
        x = this.front.Remove(this.front.Size() - i - 1)
    } else {
        x = this.back.Remove(i - this.front.Size())
    }
    this.Balance()
    return x
}

func (this *DualArrayDeque) Balance() {
    n := this.Size()
    mid := n/2
    if 3 * this.front.Size() < this.back.Size() || 3 * this.back.Size() < this.front.Size() {
        f := arraystack.New()
        for i := 0; i < mid; i++ {
            f.Add(i, this.Get(mid - i - 1))
        }
        b := arraystack.New()
        for i:= 0; i < n-mid; i++ {
            b.Add(i, this.Get(mid + i))
        }
        this.front = f
        this.back = b
    }
}