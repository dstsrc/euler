package main

import "fmt"

/*
Сумма простых чисел меньше 10 равна 2 + 3 + 5 + 7 = 17.
Найдите сумму всех простых чисел меньше двух миллионов.
*/

func main() {
	s := 0
	for i := 2; i < 2000000; i++ {
		if check(i) {
			s += i

		}
	}
	fmt.Println(s)
}
func check(d int) bool {
	for i := 2; i*i <= d; i++ {
		if d%i == 0 {
			return false
		}
	}
	return true
}
