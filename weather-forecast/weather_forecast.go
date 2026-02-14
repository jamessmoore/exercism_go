// Package weather forecasts the current 
// weather condition of various cities in Goblinocus.
package weather

var (
	// CurrentCondition represents the current weather condition.
	CurrentCondition string
	// CurrentLocation represents the current location.
	CurrentLocation  string
)


// Forecast returns the current location an current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
