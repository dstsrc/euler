package main

import "fmt"

// 137846528820
func main() {
	arr := []int{1, 1}
	cnt := 1
	for {
		if cnt == 40 {
			break
		}
		arr = step(arr)
		cnt++
	}

	fmt.Println(arr[len(arr)/2], len(arr)/2)

}

func step(arr []int) []int {
	next := make([]int, 0)
	next = append(next, 1)
	for i := 0; i < len(arr)-1; i++ {
		next = append(next, arr[i]+arr[i+1])
	}
	next = append(next, 1)
	return next
}
