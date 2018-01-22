package gointgeo

import "math"

func gcd(value1, value2 int64) int64 {
	if (value1 == 0) && (value2 == 0) {
		panic("Both values are zero.")
	}

	if value1 == math.MinInt64 {
		panic("The absolute value of the first value does not fit int64.")
	}

	if value2 == math.MinInt64 {
		panic("The absolute value of the second value does not fit int64.")
	}

	dividend := abs(value1)
	divisor := abs(value2)

	for divisor != 0 {
		mod := dividend % divisor
		dividend = divisor
		divisor = mod
	}

	return dividend
}
