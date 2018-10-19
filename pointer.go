package main

import (
    "fmt"
)

func swap1(x, y, p *int) {
    if *x > *y {
        *x, *y = *y, *x
    }
    *p = *x * *y
}

func swap2(x, y int) (int, int, int) {
    if x > y {
        x, y = y, x
    }
    return x, y, x * y
}

func main() {
    i := 9
    j := 5
    product := 0
    swap1(&i, &j, &product)
    fmt.Println(i, j, product)
    a, b, p := swap2(a, b)
    fmt.Println(a, b, p)
    c := 64
    d := 23
    c, d, f := swap2(c, d)
    fmt.Println(c, d, f)
    swap1(&c, &d, &f)
    fmt.Println(c, d, f)
}
