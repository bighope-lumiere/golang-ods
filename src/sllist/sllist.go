package sllist

type Node struct {
    x int  // データ
    next *Node  // 参照
}

func NewNode(x int) *Node {
    return &Node{
        x: x,
    }
}

type SLList struct {
    n int
    head *Node
    tail *Node
}

func New() SLList {
    return SLList{}
}

func (this *SLList) Push(x int) int {
    u := NewNode(x)
    u.next = this.head
    this.head = u
    if this.n == 0 {
        this.tail = u
    }
    this.n++
    return x
}

func (this *SLList) Pop() int {
    if this.n == 0 {
        return 0  // nil
    }
    res := this.head.x
    this.head = this.head.next
    this.n--
    if this.n == 0 {
        this.tail = nil
    }
    return res
}

func (this *SLList) Remove() int {
    return this.Pop()
}

func (this *SLList) Add(x int) bool {
    u := NewNode(x)
    if this.n == 0 {
        this.head = u
    } else {
        this.tail.next = u
    }
    this.tail = u
    this.n++
    return true
}
