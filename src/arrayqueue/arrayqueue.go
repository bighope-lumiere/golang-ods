package arrayqueue

type ArrayQueue struct {
    arr []int
    j int  // 削除する要素を追跡する index
    n int  // キューの要素数
}

func New() ArrayQueue {
    return ArrayQueue{}
}

func (this *ArrayQueue) Add(x int) bool {
    if this.n + 1 > len(this.arr) {
        this.Resize()
    }
    ind := (this.j+this.n) % len(this.arr)
    this.arr[ind] = x
    this.n++
    return true
}

func (this *ArrayQueue) Remove() int {
    x := this.arr[this.j]
    this.j = (this.j + 1) % len(this.arr)
    this.n--
    if len(this.arr) >= 3 * this.n {
        this.Resize()
    }
    return x
}

func (this *ArrayQueue) Resize() {
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