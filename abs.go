package gointgeo

import "math"

func abs(value int64) int64 {
	if value == math.MinInt64 {
		panic("The absolute value does not fit int64.")
	}

	if value < 0 {
		return -value
	} else {
		return value
	}
}
