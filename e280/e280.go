// https://projecteuler.net/problem=280
package main

import (
	"fmt"
)

const (
	GreedX = 5
	GreedY = 5
)

var ANS float64 = 0

type State struct {
	field   [GreedX][GreedY]int
	antPos  Pos
	antLoad bool
}

type Pos struct {
	X int
	Y int
}

func GetNeighPos(P Pos) []Pos {
	ans := []Pos{}
	if P.X+1 < GreedX {
		NewPos := Pos{X: P.X + 1, Y: P.Y}
		ans = append(ans, NewPos)
	}
	if P.X-1 >= 0 {
		NewPos := Pos{X: P.X - 1, Y: P.Y}
		ans = append(ans, NewPos)
	}
	if P.Y+1 < GreedY {
		NewPos := Pos{X: P.X, Y: P.Y + 1}
		ans = append(ans, NewPos)
	}
	if P.Y-1 >= 0 {
		NewPos := Pos{X: P.X, Y: P.Y - 1}
		ans = append(ans, NewPos)
	}
	return ans
}

func Init() State {
	return State{
		antLoad: false,
		antPos:  Pos{X: 2, Y: 2},
		field: [GreedX][GreedY]int{
			{-1, -1, -1, -1, -1},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1}},
	}
}

func (S *State) ChangeState() {
	if S.antPos.X == GreedX-1 && !S.antLoad && S.field[S.antPos.X][S.antPos.Y] == 1 {
		S.field[S.antPos.X][S.antPos.Y] = 0
		S.antLoad = true
	}
	if S.antPos.X == 0 && S.antLoad && S.field[S.antPos.X][S.antPos.Y] == -1 {
		S.field[S.antPos.X][S.antPos.Y] = 0
		S.antLoad = false
	}
}

func (S *State) CheckFinish() bool {
	if S.field[0] == [GreedY]int{0, 0, 0, 0, 0} {
		return true
	} else {
		return false
	}
}

func (S *State) Print() {
	fmt.Println()
	fmt.Println("greed: ")
	for i := 0; i < GreedX; i++ {
		fmt.Println(S.field[i])
	}
	fmt.Println("ant:")
	fmt.Println(S.antPos, S.antLoad)
}

func Move(S State, NextPos Pos) State {
	S.antPos = NextPos
	S.ChangeState()
	return S
}

var StateHash = make(map[State]map[int]float64)

func GetStates(S State) {
	if S.CheckFinish() {
		return
	}
	pos := GetNeighPos(S.antPos)
	for i := 0; i < len(GetNeighPos(S.antPos)); i++ {
		St := Move(S, pos[i])
		if _, ok := StateHash[St]; !ok {
			StateHash[St] = map[int]float64{}
			GetStates(St) // recursion
		}
	}
}

func (St *State) Propagate() {
	if (*St).CheckFinish() {
		for step, prob := range StateHash[*St] {
			ANS += float64(step) * prob
			delete(StateHash[*St], step)
		}
		return
	}
	neighPos := GetNeighPos(St.antPos)
	for step, prob := range StateHash[*St] {
		for i := 0; i < len(neighPos); i++ {
			nxtSt := Move(*St, neighPos[i])
			StateHash[nxtSt][step+1] += prob / float64(len(neighPos))
		}
		delete(StateHash[*St], step)
	}
}

func main() {
	St := Init()
	GetStates(St)
	StateHash[St] = map[int]float64{0: 1}

	for currStep := 0; currStep < 4300; currStep++ {
		for state, stepprob := range StateHash {
			if _, ok := stepprob[currStep]; ok {
				state.Propagate()
			}
		}
	}
	fmt.Println(ANS)
}
