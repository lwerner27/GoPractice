package main

import (
	"fmt"
)

type vehicles interface{}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type boat struct {
	vehicle
	length int
}

func main() {
	hondaCivic := car{6, 100, "blue"}
	larson := boat{vehicle{5, 65, "red"}, 15}

	// Since our "vehicles" interface is empty
	// any type we put in there will count as a vehicle
	rides := []vehicles{hondaCivic, larson}
	fmt.Println(rides)
}
