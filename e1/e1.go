package main

import "fmt"

// 233168
func main() {
	a := getSeq(1000)
	s := getSum(a)
	fmt.Println(s)
}

func getSeq(lim int) []int64 {
	s := make([]int64, 0)
	for i := 0; i < lim; i++ {
		if i%3 == 0 || i%5 == 0 {
			s = append(s, int64(i))
		}
	}
	return s
}

func getSum(a []int64) int64 {
	var s int64 = 0
	for _, v := range a {
		s += v
	}
	return s
}
