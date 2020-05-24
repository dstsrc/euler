package main

import "fmt"

/*


Простые делители числа 13195 - это 5, 7, 13 и 29.

Каков самый большой делитель числа 600851475143, являющийся простым числом?

*/

func main() {
	a := 600851475143
	m := make(map[int]int)
	getDel(a, m)
	max := getMax(m)
	fmt.Println(max)
}

func getMax(m map[int]int) int {
	var max int = 0
	for k := range m {
		if k > max {
			max = k
		}
	}
	return max
}

func getDel(a int, m map[int]int) {
	if a == 1 {
		return
	}
	for i := 2; i <= a; i++ {
		if a%i == 0 {
			m[i]++
			getDel(a/i, m)
			return
		}
	}
}
