package main

import (
	"fmt"
)

type ParkingSystem struct {
	garage [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		garage: [3]int{big, medium, small},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	this.garage[carType-1]--
	return this.garage[carType-1] >= 0
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */

func test01() {
	parking := Constructor(1, 1, 0)
	fmt.Println(parking.AddCar(1))
	fmt.Println(parking.AddCar(2))
	fmt.Println(parking.AddCar(3))
	fmt.Println(parking.AddCar(1))
}

func main() {
	test01()
}
