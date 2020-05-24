package main

import "fmt"

/*
Последовательность Фибоначчи определяется рекурсивным правилом:

    Fn = Fn−1 + Fn−2, где F1 = 1 и F2 = 1.

Таким образом, первые 12 членов последовательности равны:

    F1 = 1
    F2 = 1
    F3 = 2
    F4 = 3
    F5 = 5
    F6 = 8
    F7 = 13
    F8 = 21
    F9 = 34
    F10 = 55
    F11 = 89
    F12 = 144

Двенадцатый член F12 - первый член последовательности, который содержит три цифры.
Каков порядковый номер первого члена последовательности Фибоначчи, содержащего 1000 цифр?
*/

func main() {
	getFib()
}

func getFib() {
	a, b := []int{1}, []int{1}
	cnt := 2
	for {
		a, b = b, sum(a, b)
		cnt++
		if nums(b) > 999 {
			fmt.Println(cnt)
			return
		}
	}
}

func nums(a []int) int {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != 0 {
			return i + 1
		}
	}
	return 0
}

func sum(a, b []int) []int {
	//b > a
	m := 0

	if len(a) > len(b) {
		m = len(a)
		add := make([]int, len(a)-len(b))
		b = append(b, add...)
	} else {
		m = len(b)
		add := make([]int, len(b)-len(a))
		a = append(a, add...)
	}

	res := make([]int, 1, m)
	for i := 0; i < m; i++ {
		s := a[i] + b[i] + res[i]
		//res[i+1] = s / 10
		res = append(res, s/10)
		res[i] = s % 10
	}
	return res
}
