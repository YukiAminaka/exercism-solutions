// Package weather provides tools for weather forecasting.
package weather

// CurrentCondition is the current weather.
var CurrentCondition string
// CurrentLocation is　the current location.
var CurrentLocation string

// Forecast returns representing the weather at your current location. 
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
