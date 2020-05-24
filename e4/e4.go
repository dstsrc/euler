package main

import (
	"fmt"
	"strconv"
)

/*


Число-палиндром с обеих сторон (справа налево и слева направо) читается одинаково.
Самое большое число-палиндром, полученное умножением двух двузначных чисел – 9009 = 91 × 99.

Найдите самый большой палиндром, полученный умножением двух трехзначных чисел.

*/

func main() {
	fmt.Println(isPalindrom(1903091))
	fmt.Println(getMaxPol(1000))
}

func getMaxPol(lim int) int {
	max := 0
	for i := 0; i < lim; i++ {
		for j := 0; j < lim; j++ {
			d := i * j
			if isPalindrom(d) && d > max {
				max = d
			}
		}
	}
	return max
}

func isPalindrom(d int) bool {
	s := strconv.Itoa(d)
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
