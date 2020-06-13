package rootisharraystack

import (
    "math"
)

type RootishArrayStack struct {
    n int
    blocks [][]int  // todo ArrayStack だと []int になってしまうのでうまくいかない
}

func New() RootishArrayStack {
    return RootishArrayStack{}
}

func (this RootishArrayStack) Size() int {
    return this.n
}

func i2b(i int) int {
    return int(math.Ceil((-3.0 + math.Sqrt(float64(9 + 8 * i)))/2.0))
}

func (this RootishArrayStack) Get(i int) int {
    b := i2b(i)
    j := i - b * (b+1) / 2
    return this.blocks[b][j]
}

func (this *RootishArrayStack) Set(i, x int) int {
    b := i2b(i)
    j := i - b * (b+1) / 2
    y := this.blocks[b][j]
    this.blocks[b][j] = x
    return y
}

func (this *RootishArrayStack) Add(i, x int) {
    r := len(this.blocks)
    if r * (r+1) / 2 < this.n + 1 {
        this.Grow()
    }
    this.n++
    for j := this.n-1; j > i; j-- {
        this.Set(j, this.Get(j-1))
    }
    this.Set(i, x)
}

func (this *RootishArrayStack) Grow() {
    this.blocks = append(this.blocks, make([]int, len(this.blocks) + 1))
}

func (this *RootishArrayStack) Remove(i int) int {
    x := this.Get(i)
    for j := i; j < this.n-1; j++ {
        this.Set(j, this.Get(j+1))
    }
    this.n--
    r := len(this.blocks)
    if (r-2)*(r-1)/2 >= this.n {
        this.Shrink()
    }
    return x
}

func (this *RootishArrayStack) Shrink() {
    r := len(this.blocks)
    for r > 0 && (r-2)*(r-1)/2 >= this.n {
        ind := len(this.blocks) - 1  // remove index
        // Go では array3 := append(array1, array2...) で
        // array3 := append(array1, array2[0], array2[1], ...) みたいな感じで結合できる
        this.blocks = append(this.blocks[:ind], this.blocks[ind+1:]...)
        r--
    }
}