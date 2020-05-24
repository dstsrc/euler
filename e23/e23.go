package main

import "fmt"

//4179871
func main() {
	mIsb := make(map[int]bool)
	sum := 0
	for i := 1; i <= 28123; i++ {
		if !check(i, mIsb) {
			sum += i
		}
		if isIsb(i) {
			mIsb[i] = true
		}
	}
	fmt.Println(sum)
}

func check(i int, m map[int]bool) bool {
	for k := range m {
		if m[i-k] {
			return true
		}
	}
	return false
}
func isIsb(n int) bool {
	s := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			s += i
			if i != n/i {
				s += n / i
			}
		}
	}
	return s > n
}
