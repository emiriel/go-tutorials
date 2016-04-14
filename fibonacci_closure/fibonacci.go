package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    fibo := 0
    fiboMoins1 := 0
    return func() int {
        if fibo == 0 {
            fibo = 1
        } else {
            fibo, fiboMoins1 = fibo + fiboMoins1, fibo
        }
        return fibo
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
