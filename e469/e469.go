// https://projecteuler.net/problem=469
package main

import (
	"fmt"
	"math/big"
)

func main() {

	aux := []*big.Rat{big.NewRat(0, 1), big.NewRat(0, 1), big.NewRat(1, 1), big.NewRat(4, 3)}

	for i := 7; i <= 40; i++ {

		free := big.NewRat(0, 1)

		for j := 1; j <= i-3; j++ {

			if j == 1 || j == i-3 {
				free.Add(free, big.NewRat(1, 1))
				free.Add(free, aux[i-5])
			} else {
				free.Add(free, aux[j-2])
				free.Add(free, big.NewRat(2, 1))
				free.Add(free, aux[i-3-j-1])
			}

		}

		aux = append(aux, big.NewRat(0, 1).Mul(free, big.NewRat(1, int64(i-3))))

		ans := big.NewRat(2, 1)
		ans.Add(ans, aux[i-3])
		ans.Mul(ans, big.NewRat(1, int64(i)))

		answer, _ := ans.Float64()
		fmt.Println(i, answer)
	}

}
