package main

import (
	"fmt"
	"math"
)

//P(1, 1) = 1
//P(1, 2) = 3
//P(2, 1) = 2
//P(10, 20) = 440
//P(25, 75) = 4863
//P(99, 100) = 19454
const target = 71328803586048

func main() {
	fmt.Println(sum())
}

func getF(d2, d3 int) int {
	r := 1
	for i := 1; i <= d2; i++ {
		r = r * 2
	}
	for i := 1; i <= d3; i++ {
		r = r * 3
	}
	return r
}

//40632119
func sum() int {
	m := make(map[int]int)
	for d2 := 0; d2 <= 27; d2++ {
		for d3 := 0; d3 <= 12; d3++ {
			f := getF(d2, d3)
			r := target / f
			m[f] = r
		}
	}

	s := 0
	for f, r := range m {
		if r > 1000000000 {
			r = r % 1000000000
		}
		if f > 1000000000 {
			f = f % 1000000000
		}
		n := P(f, r)
		if n < 0 {
			panic("tut")
		}
		if n > 100000000 {
			n = n % 100000000
		}
		s += P(f, r)
		if s > 100000000 {
			s = s % 100000000
		}

	}
	return s
}

func factorize(t int) map[int]int {
	d := make(map[int]int)
	div := 2
	for t != 1 {
		if t%div == 0 {
			t = t / div
			d[div]++
		} else {
			div++
		}
	}
	return d
}

func P(f, r int) int {
	if f == 1 {
		return r * (r + 1) / 2
	}

	return (f*f)/2 + (r*(r-1))/2 + (f/2)*4*((r-f%2)/2)
}

func aux() {
	hotel := fill()
	for _, level := range hotel {
		for _, g := range level {
			fmt.Print(g, "   ")
		}
		fmt.Println()
	}
}

func fill() [][]int {
	hotel := make([][]int, 1)
	for i := 1; i <= 440; i++ {
		for j := 0; j <= 440; j++ {
			if len(hotel[j]) == 0 {
				hotel[j] = append(hotel[j], i)
				hotel = append(hotel, make([]int, 0))
				break
			} else {
				last := len(hotel[j]) - 1
				if isPerfectSquare(i + hotel[j][last]) {
					hotel[j] = append(hotel[j], i)
					break
				}
			}
		}
	}
	return hotel
}

func isPerfectSquare(x int) bool {
	r := math.Sqrt(float64(x))
	f := math.Floor(r)
	if r-f == 0 {
		return true
	}
	return false
}
