package jedlik

import "fmt"

// Add a field to the `Car` struct that keeps 
// track of the driven distance (car.go)


func (car Car) Drive() Car {
	if (car.battery >= car.batteryDrain){
		car.battery = (car.battery - car.batteryDrain)
		car.distance = (car.distance + car.speed)
		car.traveled += car.traveled
	}
	return car
}

func (car Car) DisplayDistance() Car {
	fmt.Printf("Driven %e meters", car.traveled)
	return car 
}

func (car Car) DisplayBattery() Car {
	fmt.Printf("Battery at %e%", car.battery)
	return car 
}

func (car Car) CanFinish(trackDistance int) bool {
	return false
}

