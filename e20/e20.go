package main

import "fmt"

func main() {
	d := make([]int, 1000)
	d[0] = 1
	for i := 1; i <= 100; i++ {
		d = mult(d, i)
	}
	s := 0
	for _, v := range d {
		s += v
	}
	fmt.Println(s)
}
func mult(d []int, m int) []int {
	res := make([]int, 1000)
	for i, v := range d {
		if i == len(d)-1 {
			break
		}
		t := v*m + res[i]
		res[i] = t % 10
		res[i+1] = t / 10
	}
	return res
}
