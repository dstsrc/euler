// https://projecteuler.net/problem=301
package main

import (
	"log"
	"math"
)

func main() {
	cnt := 0
	for i := 1; i <= int(math.Pow(2, 30)); i++ {
		if Check(i) == 0 {
			cnt++
		}
	}
	log.Println(cnt)

}

func Check(x int) int {
	return (x) ^ (2 * x) ^ (3 * x)
}
