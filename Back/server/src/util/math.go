package util

import "math"

const decimalPlace2Coefficient = 100

func RoundTo2DecimalPlaces(num float64) float64 {
	return math.Round(num*decimalPlace2Coefficient) / decimalPlace2Coefficient
}
