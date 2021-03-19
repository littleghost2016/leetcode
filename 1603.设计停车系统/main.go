package main

type ParkingSystem struct {
	parkingSpace [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{[3]int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.parkingSpace[carType-1] > 0 {
		this.parkingSpace[carType-1]--
		return true
	} else {
		return false
	}
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
// 每日一题，一把梭
