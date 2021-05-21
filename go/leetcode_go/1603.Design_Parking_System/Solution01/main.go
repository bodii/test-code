package main

import (
	"fmt"
)

type ParkingSystem struct {
	big    int
	medium int
	small  int
	data   []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		big:    big,
		medium: medium,
		small:  small,
		data:   []int{0, 0, 0},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.data[0] >= this.big {
			return false
		}
	case 2:
		if this.data[1] >= this.medium {
			return false
		}
	case 3:
		if this.data[2] >= this.small {
			return false
		}
	default:
		return false
	}

	this.data[carType-1]++
	return true
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
