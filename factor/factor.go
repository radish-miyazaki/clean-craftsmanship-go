package factor

func FactorsOf(n int) []int {
	var results []int

	for divisor := 2; n > 1; divisor++ {
		for n%divisor == 0 {
			results = append(results, divisor)
			n /= divisor
		}
	}

	return results
}
