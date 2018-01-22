package gointgeo

import "math"

// A fraction in its canonical representation
// with 64-bit coprime numerator and denominator
// where the denominator is positive.
type fraction64 struct {
	numerator, denominator int64
}

func newCanonicalFraction64(numerator, denominator int64) fraction64 {
	if numerator == math.MinInt64 {
		panic("The absolute value of the numerator does not fit int64.")
	}

	if denominator <= 0 {
		panic("The denominator is not positive.")
	}

	if gcd64(numerator, denominator) != 1 {
		panic("The numerator and denominator are not coprime.")
	}

	return fraction64{numerator, denominator}
}

// Creates the fraction from the numerator and the denominator
// of one of the representations of the fraction
// which does not have to be canonical.
// The numerator and the denominator do not have to be coprime.
// The denominator can be negative but cannot be zero.
// If any of the values is negative, its absolute value must fit into int64.
func newFraction64FromNonCanonical(numerator, denominator int64) fraction64 {
	if numerator == math.MinInt64 {
		panic("The absolute value of the numerator does not fit int64.")
	}

	if denominator == 0 {
		panic("The denominator is zero.")
	}

	if denominator == math.MinInt64 {
		panic("The absolute value of the denominator does not fit int64.")
	}

	fractionGcd := gcd64(numerator, denominator)

	var canonicalNumerator int64
	if denominator < 0 {
		canonicalNumerator = -numerator / fractionGcd
	} else {
		canonicalNumerator = numerator / fractionGcd
	}

	canonicalDenominator := abs64(denominator / fractionGcd)
	return newCanonicalFraction64(canonicalNumerator, canonicalDenominator)
}
