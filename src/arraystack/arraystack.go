package arraystack

type ArrayStack struct {
    n int      // number of elements
    arr []int   // array (slice) used for ArrayStack implementation
    cap int  // capacity of ArrayStack (= size of arr)
}

func New() ArrayStack {
    return ArrayStack{}
}

func (this ArrayStack) Size() int {
    return this.n
}

func (this ArrayStack) Get(i int) int {
    return this.arr[i]
}

func (this *ArrayStack) Set(i int, x int) int {
    y := this.arr[i]
    this.arr[i] = x
    return y
}

func (this *ArrayStack) Add(i int, x int) {
    if this.n == this.cap {
        this.Resize()
    }
    for j := this.n; j > i; j-- {
        this.arr[j] = this.arr[j-1]
    }
    this.arr[i] = x
    this.n++
}

func (this *ArrayStack) Remove(i int) int {
    x := this.arr[i]
    for j := i; j < this.n-1; j++ {
        this.arr[j] = this.arr[j+1]
    }
    this.n--
    if len(this.arr) >= 3 * this.n {
        this.Resize()
    }
    return x
}

func (this *ArrayStack) Resize() {
    this.cap = Max(1, 2 * this.n)
    newArray := make([]int, this.cap)
    for i := 0; i < this.n; i++ {
        newArray[i] = this.arr[i]
    }
    this.arr = newArray
}

func Max(x, y int) int {
    res := y
    if x > y {
        res = x
    }
    return res
}