package math

func MagicAdd(a, b float64) float64 {
	return a + b
}

func MagicSubtract(a, b float64) float64 {
	return a - b
}

func MagicFindMin(a, b, c int64) int64 {
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

func MagicFindMax(a, b, c int64) int64 {
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
