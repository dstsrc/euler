package main

import "fmt"

/*
Выписав первые шесть простых чисел, получим 2, 3, 5, 7, 11 и 13. Очевидно, что 6-ое простое число - 13.
Какое число является 10001-ым простым числом?
*/
//104743
func main() {
	i := 2
	cnt := 0
	for {
		if check(i) {
			//fmt.Println(i)
			cnt++
		}
		if cnt == 10001 {
			fmt.Println(i)
			return
		}
		i++
	}
}

func check(d int) bool {
	for i := 2; i < d; i++ {
		if d%i == 0 {
			return false
		}
	}
	return true
}
