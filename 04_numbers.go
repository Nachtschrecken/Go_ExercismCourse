package main

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	successRate = successRate / 100
	var calculateWorkingCarsPerHour = float64(productionRate) * successRate
	return calculateWorkingCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var calculateWorkingCarsPerMinute = CalculateWorkingCarsPerHour(productionRate, successRate) / 60
	return int(calculateWorkingCarsPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var blockOfTen = carsCount / 10
	var blockOfRest = carsCount % 10
	var calculateCost = blockOfTen*95000 + blockOfRest*10000
	return uint(calculateCost)
}
