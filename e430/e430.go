// https://projecteuler.net/problem=430
package main

import (
	"fmt"
	"math"
)

func E(N, M float64) {
	N_2 := math.Pow(N, 2)
	var i, j, ans, P_stay, P_stay_2, P_change, P_change_2, frac, elem, Sum float64
	i = 1

	for i <= N/2 {

		P_stay = (math.Pow(N-i, 2) + math.Pow(i-1, 2)) / N_2
		P_stay_2 = math.Pow(P_stay, 2)
		P_change = 1 - P_stay
		P_change_2 = math.Pow(P_change, 2)
		frac = P_change_2 / P_stay_2
		elem = math.Pow(P_stay, M)
		Sum = elem // sum for one coin
		j = 2

		for j <= M {
			elem = (elem * frac) * (M + 2 - j) / (j - 1) * (M + 1 - j) / j
			Sum += elem
			j += 2
			if elem < 0.000000000000001 {
				break
			}
		}

		ans += Sum
		i++
		if (Sum-0.5)*(N/2-i) < 0.01 {
			break
		}
		/*if int(i) % 1000 == 0 {
		    fmt.Println("op", Sum, elem)
		}*/
	}

	ans += (N/2 - i + 1) * 0.5
	fmt.Println(2 * ans)
	fmt.Printf("%.2f\n", 2*ans)
}

func main() {
	E(10000000000, 4000)
}

//  5000624921.38
