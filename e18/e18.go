package main

import (
	"fmt"
	"strconv"
	"strings"
)

const strPiram = `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

const strPiramTest = `3
7 4
2 4 6
8 5 9 3`

func main() {
	piram := getPiram(strPiram)
	fmt.Println(piram)
	ans := proc(piram)
	fmt.Println(ans)
}

func proc(p [][]int) int {
	for i := len(p) - 1; i > 0; i-- {

		for j := len(p[i]) - 1; j > 0; j-- {
			s1 := p[i][j] + p[i-1][j-1]
			s2 := p[i][j-1] + p[i-1][j-1]
			if s1 > s2 {
				p[i-1][j-1] = s1
			} else {
				p[i-1][j-1] = s2
			}
		}

	}
	return p[0][0]
}

func getPiram(s string) [][]int {
	piram := make([][]int, 0)

	strRows := strings.Split(s, "\n")

	for _, row := range strRows {
		r := make([]int, 0)
		for _, item := range strings.Split(row, " ") {
			d, err := strconv.Atoi(item)
			if err == nil {
				r = append(r, d)
			}
		}
		piram = append(piram, r)

	}

	return piram
}
