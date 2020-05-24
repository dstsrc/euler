package main

import "fmt"

/*
2520 - самое маленькое число, которое делится без остатка на все числа от 1 до 10.
Какое самое маленькое число делится нацело на все числа от 1 до 20?
*/
//232792560
func main() {
	i := 1
	for {
		if check(i, 20) {
			fmt.Println(i)
			break
		}
		i++
	}
}

func check(d, n int) bool {
	for i := 1; i <= n; i++ {
		if d%i != 0 {
			return false
		}
	}
	return true
}
