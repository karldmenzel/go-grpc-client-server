package math

func LocalAdd(a, b float64) float64 {
	return a + b
}

func LocalSubtract(a, b float64) float64 {
	return a - b
}

func LocalFindMin(a, b, c int64) int64 {
	if (a < b) && (a < c) {
		return a
	}

	if (b < a) && (b < c) {
		return b
	}

	if (c < a) && (c < b) {
		return c
	}

	return a
}

func LocalFindMax(a, b, c int64) int64 {
	if (a > b) && (a > c) {
		return a
	}

	if (b > a) && (b > c) {
		return b
	}

	if (c > a) && (c > b) {
		return c
	}

	return a
}
