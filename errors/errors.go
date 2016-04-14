package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64 

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Erreur : la racine carrée de %v n'existe pas dans l'ensemble des réels", float64(e));
}

const (
    NEWTON_DELTA = 0.00000000001
)

func between(aComparer, min, max float64) (resultat bool) {
    resultat = aComparer >= min && aComparer <= max
    return
}

func Sqrt(x float64) (float64, error) {
	if(x < 0) {
		return 0, ErrNegativeSqrt(x)
	}
	var res float64 = 1.0
	var previousRes float64 = 0.0
	for diff := -1.0 ; !between(diff, -NEWTON_DELTA, NEWTON_DELTA) ; diff = (res - previousRes) / res {
		previousRes = res
		res = previousRes - ((math.Pow(previousRes, 2) - x) / (2 * previousRes))
	}
	return res, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
