package main

import "fmt"

type vehicles interface{} //empty interface

type isplane interface {
	wingspan() float64
}

type vehicle struct {
	TopSpeed int
	Seats    int
	Color    string
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

func (p plane) wingspan() float64 {
	if p.Jet {
		return float64(float64(p.Seats) * 1.4)
	} else {
		return float64(float64(p.Seats) * 3.2)
	}
}

func calculatefuel(p isplane) float64 {
	return p.wingspan() * 4
}

func details(p vehicles) {
	fmt.Println(p)
}

func main() {
	bavaria := boat{vehicle{22, 4, "White"}, 34} //A boat with a max speed of 22, 4 seats, White colour, and 34 feet
	car := car{vehicle{90, 5, "Red"}, 4, 5}
	a757 := plane{vehicle{400, 100, "Grey"}, true}

	calculatefuel(a757)
	details(bavaria)
	details(car)
	rides := []vehicles{bavaria, car, a757}
	fmt.Println(rides)

}
