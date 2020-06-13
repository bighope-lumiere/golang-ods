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

// Exercise 2.1: AddAll(i, c)
func (this *ArrayStack) AddAll(i int, c ArrayStack) {
    m := c.Size()
    // 配列のサイズが足りない場合の処理（resizeに相当）
    if this.n + m > this.cap {
        newCap := Max(1, this.n)
        for newCap < this.n + m {
            newCap *= 2
        }
        this.cap = newCap
        newArray := make([]int, this.cap)
        for j := 0; j < this.n; j++ {
            newArray[j] = this.arr[j]
        }
        this.arr = newArray
    }

    // i番目以降を m だけ後ろにずらす
    for j := this.n-1; j >= i; j-- {
        this.arr[j+m] = this.arr[j]
    }

    // c の要素を1つずつ入れる
    for j := 0; j < m; j++ {
        this.arr[i+j] = c.Get(j)
    }
    this.n += m
}