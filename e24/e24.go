package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 2; i <= 1000000; i++ {
		a = getNExt(a)
	}
	fmt.Println(a)
}

func getNExt(a []int) []int {
	for i := len(a) - 1; i > 0; i-- {
		if a[i] > a[i-1] {
			j := find(a, a[i-1])
			a[i-1], a[j] = a[j], a[i-1]
			revers(a[i:])
			break
		}
	}
	return a
}
func revers(a []int) {
	for j, k := 0, len(a)-1; j < k; j, k = j+1, k-1 {
		a[j], a[k] = a[k], a[j]
	}
}

func find(a []int, p int) int {
	for i := len(a) - 1; i > 0; i-- {
		if a[i] > p {
			return i
		}
	}
	return 0
}
