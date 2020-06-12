package main
import (
    "fmt"
    "arraystack"
)

func main() {
    N := 100
    as := arraystack.New()
    fmt.Println("making ArrayStack as")
    for i := 0; i < N; i++ {
        as.Add(i, i)
    }
    fmt.Println("adding as done")
    as.Set(88, 199)
    as.Add(87, 133)
    fmt.Println(as.Remove(98))
    for i := 0; i < 80; i++ {
        fmt.Println("removed = ", as.Remove(0))
    }
    for i := 0; i < as.Size(); i++ {
        fmt.Println("as.get(", i, ")=", as.Get(i))
    }
}