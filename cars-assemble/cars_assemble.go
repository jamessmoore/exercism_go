// Package cars calculates the number of 
// working cars produced per hour
package cars

// productionRate represents
var productionRate int = 1547

// successRate represents 
var successRate float64 = 90

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int((float64(productionRate) / 60) * (successRate / 100))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var cost uint = 0
	groups := carsCount / 10
	remCars := carsCount % 10
	cost = (uint(groups) * 95000) + (uint(remCars) * 10000)
	return cost
}