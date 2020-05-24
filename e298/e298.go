// https://projecteuler.net/problem=298
package main

import (
	"fmt"
	"math"
)

type UniqState struct {
	MemR [5]string //{"A","B","C","D","E"}
	MemL [5]string
}

type State struct {
	MemR [5]int
	MemL [5]int
}

func Step(St State, Num int) (State, int) {
	// arrival no change order for R:
	var ScoreR int = 0
	var ZeroIndexR int = -1
	for i, v := range St.MemR {
		if v == Num {
			ScoreR++
		}
		if v == 0 && ZeroIndexR == -1 {
			ZeroIndexR = i
		}
	}
	if ScoreR == 0 {
		if ZeroIndexR != -1 {
			St.MemR[ZeroIndexR] = Num
		} else {
			//tmp := Num
			for i := range St.MemR {
				if i == len(St.MemR)-1 {
					St.MemR[i] = Num
				} else {
					St.MemR[i] = St.MemR[i+1]
				}
			}
		}
	}
	// mention change order for L:
	var ScoreL int = 0
	var ZeroIndexL int = -1
	var Hit int = -1
	for i, v := range St.MemL {
		if v == Num {
			ScoreL++
			Hit = i
		}
		if v == 0 && ZeroIndexL == -1 {
			ZeroIndexL = i
		}
	}
	if ScoreL == 0 {
		if ZeroIndexL != -1 {
			St.MemL[ZeroIndexL] = Num
		} else {
			//tmp := Num
			for i := range St.MemL {
				if i == len(St.MemL)-1 {
					St.MemL[i] = Num
				} else {
					St.MemL[i] = St.MemL[i+1]
				}
			}
		}
	} else {
		var stop int
		if ZeroIndexL != -1 {
			stop = ZeroIndexL
		} else {
			stop = len(St.MemL)
		}
		for i := Hit; i < stop; i++ {
			if i == stop-1 {
				St.MemL[i] = Num
			} else {
				St.MemL[i] = St.MemL[i+1]
			}

		}
	}
	return St, ScoreR - ScoreL
}

func (St State) StToUSt() (USt UniqState) {
	Letters := [5]string{"A", "B", "C", "D", "E"}
	DifLetters := [5]string{"F", "G", "H", "I", "J"}
	for i, vi := range St.MemR {
		if vi == 0 {
			break
		} else {
			USt.MemR[i] = Letters[i]
			for j, vj := range St.MemL {
				if vi == vj {
					USt.MemL[j] = Letters[i]
				}
			}
		}
	}
	tmp := 0
	for i, v := range USt.MemL {
		if tmp == 3 {
			fmt.Println("blyat problem")
		}
		if v == "" && St.MemL[i] != 0 {
			USt.MemL[i] = DifLetters[tmp]
			tmp++

		}
	}
	return USt
}

func GetNum(Str string) (Num int) {
	var ans int
	switch Str {
	case "A":
		ans = 1
	case "B":
		ans = 2
	case "C":
		ans = 3
	case "D":
		ans = 4
	case "E":
		ans = 5
	case "F":
		ans = 6
	case "G":
		ans = 7
	case "H":
		ans = 8
	case "I":
		ans = 9
	case "":
		ans = 0
	}
	return ans
}

func (USt UniqState) UStToSt() (St State) {
	for i, v := range USt.MemL {
		St.MemL[i] = GetNum(v)
	}
	for i, v := range USt.MemR {
		St.MemR[i] = GetNum(v)
	}
	return St
}

func ForWard(StepHash map[UniqState]map[int]float64) map[UniqState]map[int]float64 {
	NextStepHash := make(map[UniqState]map[int]float64)
	for Ust, DiffProb := range StepHash {
		for Diff, Prob := range DiffProb {
			st := Ust.UStToSt()

			for i := 1; i <= 10; i++ {
				stm, diff := Step(st, i)
				ust := stm.StToUSt()
				if _, ok1 := NextStepHash[ust]; ok1 {

					if _, ok2 := NextStepHash[ust][Diff+diff]; ok2 {
						NextStepHash[ust][Diff+diff] += 0.1 * Prob
					} else {
						NextStepHash[ust][Diff+diff] = 0.1 * Prob
					}

				} else {
					NextStepHash[ust] = map[int]float64{Diff + diff: 0.1 * Prob}
				}
			}
		}
	}
	return NextStepHash
}

func main() {
	// init
	StepHash := make(map[UniqState]map[int]float64)
	var st State
	for i := 1; i <= 10; i++ {
		stm, diff := Step(st, i)
		ust := stm.StToUSt()
		prob := StepHash[ust][diff]
		StepHash[ust] = map[int]float64{diff: prob + 0.1}
	}

	for i := 2; i <= 50; i++ {
		//fmt.Println(i)
		StepHash = ForWard(StepHash)
	}

	var ans float64
	for _, DiffProb := range StepHash {
		for Diff, Prob := range DiffProb {
			if Diff != 0 {
				ans += (math.Abs(float64(Diff)) * Prob)
			}
		}
	}
	fmt.Println(ans)
}

/*   old code based on rand
package main

import (
	"fmt"
	"github.com/bradfitz/slice"
	"math"
	"math/rand"
	"time"
)

type CellMemL struct {
	number      int
	mentionStep int
}

type CellMemR struct {
	number      int
	arrivalStep int
}

type TypeRobin struct {
	Memory [5]CellMemR
	Score  int
}

type TypeLarry struct {
	Memory [5]CellMemL
	Score  int
}

func (Larry *TypeLarry) LarryGetNumber(Num, Step int) {
	var lastMentionPos int = 0
	update := true
	for i := range Larry.Memory {
		if Larry.Memory[i].number == Num {
			Larry.Score++
			Larry.Memory[i].mentionStep = Step
			update = false
			break
		}
		if Larry.Memory[i].mentionStep < Larry.Memory[lastMentionPos].mentionStep {
			lastMentionPos = i
		}
	}
	if update {
		Larry.Memory[lastMentionPos] = CellMemL{number: Num, mentionStep: Step}
	}
}

func (Robin *TypeRobin) RobinGetNumber(Num, Step int) {
	var arrivalStepPos int = 0
	update := true
	for i := range Robin.Memory {
		if Robin.Memory[i].number == Num {
			Robin.Score++
			update = false
			break
		}
		if Robin.Memory[i].arrivalStep < Robin.Memory[arrivalStepPos].arrivalStep {
			arrivalStepPos = i
		}
	}
	if update {
		Robin.Memory[arrivalStepPos] = CellMemR{number: Num, arrivalStep: Step}
	}
}

func epoch(steps int) (float64, float64, float64) {
	var Larry TypeLarry
	var Robin TypeRobin
	endStep := steps
	for step := 1; step <= endStep; step++ {
		num := rand.Intn(10) + 1
		Robin.RobinGetNumber(num, step)
		Larry.LarryGetNumber(num, step)

	}
	return float64(Larry.Score), float64(Robin.Score), math.Abs(float64(Robin.Score - Larry.Score))
}

func SupportProcess(steps, epochNum int) {
	var sum float64 = 0
	for i := 1; i <= epochNum; i++ {
		_, _, tmp := epoch(steps)
		sum += tmp
	}
	fmt.Println(sum / float64(epochNum))
}

func AVGLarry(steps, epochNum int) float64 {
	var sum float64 = 0
	for i := 1; i <= epochNum; i++ {
		tmp, _, _ := epoch(steps)
		sum += tmp
	}
	return (float64(sum) / float64(epochNum))
}

func AVGRobin(steps, epochNum int) float64 {
	var sum float64 = 0
	for i := 1; i <= epochNum; i++ {
		_, tmp, _ := epoch(steps)
		sum += tmp
	}
	return (float64(sum) / float64(epochNum))
}

func SupportProcess2(set []int) float64 {
	var Larry TypeLarry
	var Robin TypeRobin
	for step := 1; step < len(set); step++ {
		Robin.RobinGetNumber(set[step], step)
		Larry.LarryGetNumber(set[step], step)
	}

	slice.Sort(Larry.Memory[:], func(i, j int) bool { return Larry.Memory[i].mentionStep < Larry.Memory[j].mentionStep })
	slice.Sort(Robin.Memory[:], func(i, j int) bool { return Robin.Memory[i].arrivalStep < Robin.Memory[j].arrivalStep })

	fmt.Println("mention", Larry)
	fmt.Println("arrival", Robin)
	fmt.Println(math.Abs(float64(Robin.Score - Larry.Score)))
	return math.Abs(float64(Robin.Score - Larry.Score))
}

// on 8 step 1st difference 0.01208765
// 4 (var of mention) 0.1 (prob of mention) * 1 * 0.9 * 0.8 * 0.7 * 0.6 (first 5 pos) * 0.5 (new elem) * 0.2 (first or second item) =
// = 0.012096

//on 9 step difference 0.044758381 0.04473222 0.04478199 0.04474649
// 8 0.012096
// 9 0.0447087
// 10 0.0987495
// 11 0.1688996
// 12 0.248784
// 13 0.3312182
// 14 0.4125293
// 15 0.4896989
// 16 0.5609053
// 20 0.799171
// 30 1.2127731
// 40 1.5174004
// 50 1.7688169
// 80 2.3670867
// 110 2.8409253
// 140 3.2442027
// 200 3.9322287
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	//fmt.Println(4 * 0.1 * 1 * 0.9 * 0.8 * 0.7 * 0.6 * 0.5 * 0.2)
	SupportProcess(200, 10000)
	//SupportProcess2([]int{0, 1, 1})
	//fmt.Println(AVGLarry(4, 100000000))
	//fmt.Println(AVGRobin(4, 100000000))
}

*/
