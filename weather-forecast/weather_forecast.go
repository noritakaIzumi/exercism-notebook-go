// Package weather
// forecast package.
package weather

// CurrentCondition is a current weather.
var CurrentCondition string

// CurrentLocation is a current location such as city.
var CurrentLocation string

// Forecast returns a formatted string of weather forecast.
//goland:noinspection GoUnusedExportedFunction
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
