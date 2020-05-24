package main

import "fmt"

/*
Следующая повторяющаяся последовательность определена для множества натуральных чисел:

n → n/2 (n - четное)
n → 3n + 1 (n - нечетное)

Используя описанное выше правило и начиная с 13, сгенерируется следующая последовательность:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

Получившаяся последовательность (начиная с 13 и заканчивая 1) содержит 10 элементов. Хотя это до сих пор и не доказано (проблема Коллатца (Collatz)), предполагается, что все сгенерированные таким образом последовательности оканчиваются на 1.

Какой начальный элемент меньше миллиона генерирует самую длинную последовательность?

Примечание: Следующие за первым элементы последовательности могут быть больше миллиона.
*/
func main() {
	s := proc()
	fmt.Println(s)
}

func proc() int {
	m := map[int]int{8: 4}
	max := 0
	start := 0
	for i := 2; i < 1000000; i++ {
		cnt := getLenSeq(i, m)
		m[i] = cnt
		if cnt > max {
			max = cnt
			start = i
		}
	}
	return start
}

func getLenSeq(n int, m map[int]int) int {
	cnt := 1
	for {
		if v, ok := m[n]; ok {
			return cnt + v - 1
		}
		n = getNext(n)
		cnt++
		if n == 1 {
			break
		}
	}
	return cnt
}

func getNext(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}
