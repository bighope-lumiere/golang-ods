package arraystack

type ArrayStack struct {
    n int      // number of elements
    array []int   // array (slice) used for ArrayStack implementation
    cap int  // capacity of ArrayStack (= size of array)
}

func New() ArrayStack {
    return ArrayStack{}
}

func (AS ArrayStack) Size() int {
    return AS.n
}

func (AS ArrayStack) Get(i int) int {
    return AS.array[i]
}

func (AS *ArrayStack) Set(i int, x int) int {
    y := AS.array[i]
    AS.array[i] = x
    return y
}

func (AS *ArrayStack) Add(i int, x int) {
    if AS.n == AS.cap {
        AS.Resize()
    }
    for j := AS.n; j > i; j-- {
        AS.array[j] = AS.array[j-1]
    }
    AS.array[i] = x
    AS.n++
}

func (AS *ArrayStack) Remove(i int) int {
    x := AS.array[i]
    for j := i; j < AS.n-1; j++ {
        AS.array[j] = AS.array[j+1]
    }
    AS.n--
    if len(AS.array) >= 3 * AS.n {
        AS.Resize()
    }
    return x
}

func (AS *ArrayStack) Resize() {
    AS.cap = Max(1, 2 * AS.n)
    newArray := make([]int, AS.cap)
    for i := 0; i < AS.n; i++ {
        newArray[i] = AS.array[i]
    }
    AS.array = newArray
}

func Max(x, y int) int {
    res := y
    if x > y {
        res = x
    }
    return res
}