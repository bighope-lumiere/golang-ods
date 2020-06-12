package arraystack

type ArrayStack struct {
    n int      // number of elements
    array []int   // array (slice) used for ArrayStack implementation
    cap int  // capacity of ArrayStack (= size of array)
}

func New() ArrayStack {
    return ArrayStack{}
}

func (this ArrayStack) Size() int {
    return this.n
}

func (this ArrayStack) Get(i int) int {
    return this.array[i]
}

func (this *ArrayStack) Set(i int, x int) int {
    y := this.array[i]
    this.array[i] = x
    return y
}

func (this *ArrayStack) Add(i int, x int) {
    if this.n == this.cap {
        this.Resize()
    }
    for j := this.n; j > i; j-- {
        this.array[j] = this.array[j-1]
    }
    this.array[i] = x
    this.n++
}

func (this *ArrayStack) Remove(i int) int {
    x := this.array[i]
    for j := i; j < this.n-1; j++ {
        this.array[j] = this.array[j+1]
    }
    this.n--
    if len(this.array) >= 3 * this.n {
        this.Resize()
    }
    return x
}

func (this *ArrayStack) Resize() {
    this.cap = Max(1, 2 * this.n)
    newArray := make([]int, this.cap)
    for i := 0; i < this.n; i++ {
        newArray[i] = this.array[i]
    }
    this.array = newArray
}

func Max(x, y int) int {
    res := y
    if x > y {
        res = x
    }
    return res
}