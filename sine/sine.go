package sine

import "math"

func Sin(rad float64) float64 {
	rad = math.Mod(rad, 2*math.Pi)
	result := rad
	lastResult := 2.0
	m1 := -1.0
	sign := 1.0
	power := rad
	var fac float64 = 1
	r2 := rad * rad
	n := 1
	for !closeCondition(result, lastResult) {
		lastResult = result
		power *= r2
		fac *= float64((n + 1) * (n + 2))
		n += 2
		sign *= m1
		term := sign * power / fac
		result += term
	}

	return result
}

func closeCondition(a float64, b float64) bool {
	return math.Abs(a-b) < .0000001
}
