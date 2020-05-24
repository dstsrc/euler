// https://projecteuler.net/problem=460
package main

import (
	"fmt"
	"math"
)

//F(4) ≈ 2.960516287 (0,1)->(1,2)->(2,2)
//F(10) ≈ 4.668187834 {0 1} {1 3} {2 4} {4 5} {5 5}
//F(100) ≈ 9.217221972
//F(200) 10.600176620292238
//F(400) 11.984555045810415
//F(800) 13.370040568469983
//F(1000) 13.816157113969712
//F(10000) 18.420738198997643

type Point struct {
	X float64
	Y float64
}

type Data struct {
	Time float64
	Prev Point
}

type Truck []Point

func main() {
	//SupportProcess(100)

	Process(10000)
}

func Process(steps int) {
	R := float64(steps / 2)
	MyTrack := GetTrack(R)
	var TimeHash = make(map[Point]float64)
	TimeHash[MyTrack[0]] = 0
	for i := 0; i < len(MyTrack); i++ {
		for j := 0; j < i; j++ {
			DTime := GetTimeFromTo(MyTrack[j], MyTrack[i])
			if T, ok := TimeHash[MyTrack[i]]; ok && T > TimeHash[MyTrack[j]]+DTime {
				TimeHash[MyTrack[i]] = TimeHash[MyTrack[j]] + DTime
			} else if !ok {
				TimeHash[MyTrack[i]] = DTime
			}
		}
	}
	fmt.Println(2 * TimeHash[Point{R, R}])
}

func SupportProcess(steps int) { //only for anderstand what ants track is half circle !
	var DataHash = make(map[Point]Data)
	P0 := Point{0, 1}
	DataHash[P0] = Data{0, P0}

	D := steps / 2
	for X := 1; X <= D; X++ {
		Y := X + 1

		for i := 0; i < X; i++ {
			for j := i; j < Y; j++ { // in the middle of 45 and 90 degrees
				P0 = Point{float64(i), float64(j)}

				for x := i; x <= X; x++ { // top line of new points
					tmpP := Point{float64(x), float64(Y)}
					tmpTime := GetTimeFromTo(P0, tmpP)
					if T, ok := DataHash[tmpP]; ok && T.Time > DataHash[P0].Time+tmpTime {
						DataHash[tmpP] = Data{tmpTime + DataHash[P0].Time, P0}
					} else if !ok {
						DataHash[tmpP] = Data{tmpTime + DataHash[P0].Time, P0}
					}
				}
				for y := Y - 1; y <= Y; y++ { //only 2 point
					tmpP := Point{float64(X), float64(y)}
					tmpTime := GetTimeFromTo(P0, tmpP)
					if T, ok := DataHash[tmpP]; ok && T.Time > DataHash[P0].Time+tmpTime {
						DataHash[tmpP] = Data{tmpTime + DataHash[P0].Time, P0}
					} else if !ok {
						DataHash[tmpP] = Data{tmpTime + DataHash[P0].Time, P0}
					}
				}

			}
		}

	}

	fmt.Println(steps, 2*DataHash[Point{float64(steps / 2), float64(steps / 2)}].Time)
	Pnt := Point{float64(steps / 2), float64(steps / 2)}
	SP := Point{0, 1}
	for Pnt != SP {
		if (float64(D)-Pnt.X)/Pnt.Y > 1 { // near start point
			fmt.Println(Pnt, GetX(Pnt.Y, float64(D)), Pnt.Y)
		} else { // near end point
			fmt.Println(Pnt, Pnt.X, GetY(Pnt.X, float64(D)))
		}

		Pnt = DataHash[Pnt].Prev
	}
	fmt.Println(SP, GetX(SP.Y, float64(D)), SP.Y)
	fmt.Println("-------------------------")
}

func GetX(Y, R float64) float64 {
	return math.Floor(R - math.Sqrt(R*R-(Y)*(Y)) + 0.5)
}

func GetY(X, R float64) float64 {
	return math.Floor(math.Sqrt(R*R-(X-R)*(X-R)) + 0.5)
}

func GetTrack(R float64) Truck {
	var MyTruck Truck
	MyTruck = append(MyTruck, Point{0, 1})
	check := true
	y := 2
	for check {
		if check {
			P := Point{GetX(float64(y), R), float64(y)}
			if (R-P.X)/P.Y >= 1 {
				MyTruck = append(MyTruck, P)
			} else {
				check = false
			}
		}
		y++
	}
	for x := int(MyTruck[len(MyTruck)-1].X) + 1; x <= int(R); x++ {
		MyTruck = append(MyTruck, Point{float64(x), GetY(float64(x), R)})
	}
	return MyTruck
}

func GetTimeFromTo(P0, P1 Point) float64 {
	var V float64
	if P1.Y == P0.Y {
		V = P1.Y
	} else {
		V = (P1.Y - P0.Y) / (math.Log(P1.Y / P0.Y))
	}
	R := math.Sqrt((P1.X-P0.X)*(P1.X-P0.X) + (P1.Y-P0.Y)*(P1.Y-P0.Y))
	return R / V
}
