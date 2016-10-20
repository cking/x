package xmath

// Round Round a number to its nearest int
func Round(f float64) int64 {
	if f < -0.5 {
		return int64(f - 0.5)
	}

	if f > 0.5 {
		return int64(f + 0.5)
	}

	return 0
}
