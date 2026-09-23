package util

import "math"

/* Determine if the specified year is a leap year. */
func IsLeapYear(inputYear int) bool {
	var year float64 = float64(inputYear)

	if int(year)%4 == 0 {
		if int(year)%100 == 0 {
			if int(year)%400 == 0 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	} else {
		return false
	}
}

/* Round a float to the specified number of decimal places. */
func RoundTo(inputValue float64, decimalPlaces int) float64 {
	var p float64 = math.Pow10(decimalPlaces)

	return math.Round(inputValue*p) / p
}

/* Convert degrees to radians */
func DegreesToRadians(degrees float64) float64 {
	return (degrees * math.Pi) / 180
}

/* Convert radians to degrees */
func RadiansToDegrees(radians float64) float64 {
	return (radians * 180) / math.Pi
}
