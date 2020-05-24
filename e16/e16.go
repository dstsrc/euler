package main

import "fmt"

/*
215 = 32768, сумма цифр этого числа равна 3 + 2 + 7 + 6 + 8 = 26.

Какова сумма цифр числа 21000?

*/

func main() {
	d := make([]int, 1000)
	d[0] = 1
	for i := 0; i < 1000; i++ {
		d = mult(d, 2)
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
