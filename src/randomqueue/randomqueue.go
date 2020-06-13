package randomqueue

import (
    "math/rand"
)

type RandomQueue struct {
    arr []int
    j int  // 削除する要素を追跡する index
    n int  // キューの要素数
}

func New() RandomQueue {
    return RandomQueue{}
}

func (this RandomQueue) Size() int {
    return this.n
}

func (this RandomQueue) Get(i int) int {
    return this.arr[i]
}

func (this *RandomQueue) Set(i int, x int) int {
    y := this.arr[i]
    this.arr[i] = x
    return y
}

func (this *RandomQueue) Add(x int) bool {
    if this.n + 1 > len(this.arr) {
        this.Resize()
    }
    ind := (this.j+this.n) % len(this.arr)
    this.arr[ind] = x
    this.n++
    return true
}

func (this *RandomQueue) Remove() int {
    last_ind := this.Size() - 1;
    remove_ind := rand.Intn(this.Size())
    tmp := this.Get(last_ind)
    this.Set(last_ind, this.Get(remove_ind))
    this.Set(remove_ind, tmp)

    // 以下が ArrayQueue の Remove()
    x := this.arr[this.j]
    this.j = (this.j + 1) % len(this.arr)
    this.n--
    if len(this.arr) >= 3 * this.n {
        this.Resize()
    }
    return x
}

func (this *RandomQueue) Resize() {
    newArray := make([]int, Max(1, 2 * this.n))
    for k := 0; k < this.n; k++ {
        newArray[k] = this.arr[(this.j + k) % len(this.arr)]
    }
    this.arr = newArray
    this.j = 0
}

func Max(x, y int) int {
    res := y
    if x > y {
        res = x
    }
    return res
}