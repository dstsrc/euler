package main

import "fmt"

/*
Пусть d(n) определяется как сумма делителей n (числа меньше n, делящие n нацело).
Если d(a) = b и d(b) = a, где a ≠ b, то a и b называются дружественной парой, а каждое из чисел a и b - дружественным числом.

Например, делителями числа 220 являются 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 и 110,
поэтому d(220) = 284. Делители 284 - 1, 2, 4, 71, 142, поэтому d(284) = 220.

Подсчитайте сумму всех дружественных чисел меньше 10000.

*/
//31626
func main() {
	m := make(map[int]bool)
	for i := 2; i < 10000; i++ {
		//if m[i] {
		//	//continue
		//}
		m = frend(i, m)
	}

	s := 0
	for k := range m {
		s += k
	}
	fmt.Println(m)
	fmt.Println(s)
}

func frend(n int, m map[int]bool) map[int]bool {
	f := d(n)
	if d(f) == n && n != f {
		m[n] = true
		m[f] = true
	}
	return m
}

func d(n int) int {
	s := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			s += i
			if i != n/i {
				s += n / i
			}
		}
	}
	return s
}
