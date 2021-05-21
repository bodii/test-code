package main

import (
	"fmt"
)

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		big:    big,
		medium: medium,
		small:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if carType == 1 && this.big > 0 {
		this.big--
		return true
	} else if carType == 2 && this.medium > 0 {
		this.medium--
		return true
	} else if carType == 3 && this.small > 0 {
		this.small--
		return true
	}

	return false
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
