package arraydeque

type ArrayDeque struct {
    arr []int
    j int
    n int
}

func New() ArrayDeque {
    return ArrayDeque{}
}

func (this ArrayDeque) Get(i int) int {
    return this.arr[(i + this.j) % len(this.arr)]
}

func (this *ArrayDeque) Set(i, x int) int {
    y := this.arr[(i + this.j) % len(this.arr)]
    this.arr[(i + this.j) % len(this.arr)] = x
    return y
}

func (this *ArrayDeque) Add(i, x int) {
    if this.n == len(this.arr) {
        this.Resize()
    }

    if i < this.n/2 {
        this.j = (this.j - 1) % len(this.arr)
        for k := 0; k < i; k++ {
            this.arr[(this.j + k) % len(this.arr)] = this.arr[(this.j + k + 1) % len(this.arr)]
        }
    } else {
        for k := this.n; k > i; k-- {
            this.arr[(this.j + k) % len(this.arr)] = this.arr[(this.j + k - 1) % len(this.arr)]
        }
    }

    this.arr[(this.j + i) % len(this.arr)] = x
    this.n++
}

func (this *ArrayDeque) Remove(i int) int {
    x := this.arr[(this.j + i) % len(this.arr)]

    if i < this.n/2 {
        for k := i; k > 0; k-- {
            this.arr[(this.j + k) % len(this.arr)] = this.arr[(this.j + k - 1) % len(this.arr)]
        }
        this.j = (this.j + 1) % len(this.arr)
    } else {
        for k := i; k < this.n-1; k++ {
            this.arr[(this.j + k) % len(this.arr)] = this.arr[(this.j + k + 1) % len(this.arr)]
        }
    }

    this.n--
    if len(this.arr) >= 3 * this.n {
        this.Resize()
    }
    return x
}

func (this *ArrayDeque) Resize() {
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