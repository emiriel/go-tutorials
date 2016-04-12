package main

import (
    "fmt"
    "math"
)

const (
    NEWTON_DELTA = 0.00000000001
)

func between(aComparer, min, max float64) (resultat bool) {
    resultat = aComparer >= min && aComparer <= max
    return
}

func Sqrt(x float64) float64 {
    var res float64 = 1.0
    var previousRes float64 = 0.0
    for diff := -1.0 ; !between(diff, -NEWTON_DELTA, NEWTON_DELTA) ; diff = (res - previousRes) / res {
            previousRes = res
            res = previousRes - ((math.Pow(previousRes, 2) - x) / (2 * previousRes))
    }
    return res
}

func main() {
    fmt.Println(Sqrt(7))
    fmt.Println(math.Sqrt(7))
}
