package math32

// Tolerance approximates equality between a and b within (absolute) epsilon.
func Tolerance(a, b, epsilon float32) bool {
	d := a - b
	if d < 0 {
		d = -d
	}

	if b != 0 {
		epsilon *= b
		if epsilon < 0 {
			epsilon = -epsilon
		}
	}

	return d < epsilon
}

// Close approximates equality between a and b with 1e-6 precision.
func Close(a, b float32) bool {
	return Tolerance(a, b, 1e-6)
}
