package main

import "fmt"

/*

Каждый следующий элемент ряда Фибоначчи получается при сложении двух предыдущих. Начиная с 1 и 2, первые 10 элементов будут:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

Найдите сумму всех четных элементов ряда Фибоначчи, которые не превышают четыре миллиона.
*/
// 4613732
func main() {
	seq := getFib()
	sum := getSum(seq)
	fmt.Println(sum)
}

func getSum(a []int64) int64 {
	s := int64(0)
	for _, v := range a {
		if v%2 == 0 {
			s += v
		}
	}
	return s
}

func getFib() []int64 {
	seq := make([]int64, 0)
	a, b := int64(0), int64(1)
	for b < 4000000 {
		a, b = b, a+b
		seq = append(seq, b)
	}
	return seq
}
