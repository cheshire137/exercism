package cars

import "math"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / float64(100))
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    carsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
    carsPerMin := carsPerHour / float64(60)
    return int(math.Floor(carsPerMin))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    remainder := carsCount % 10
    groupsOfTen := (carsCount - remainder) / 10
	return uint((groupsOfTen * 95000) + (remainder * 10000))
}
