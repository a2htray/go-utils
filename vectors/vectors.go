package vectors


/* 向量的四则运算 */
// Add ...
func Add(a, b []float64) []float64 {
	n1 := len(a)
	n2 := len(b)
	var n int
	switch n1 > n2 {
	case true:
		n = n2
	default:
		n = n1
	}

	ret := make([]float64, n)
	for i := 0; i < n; i++ {
		ret[i] = a[i] + b[i]
	}
	return ret
}

// Sub ...
func Sub(a, b []float64) []float64 {
	n1 := len(a)
	n2 := len(b)
	var n int
	switch n1 > n2 {
	case true:
		n = n2
	default:
		n = n1
	}

	ret := make([]float64, n)
	for i := 0; i < n; i++ {
		ret[i] = a[i] - b[i]
	}
	return ret
}

// Multiply ...
func Multiply(a []float64, factor float64) []float64 {
	ret := make([]float64, len(a))
	copy(ret, a)

	for i, v := range ret {
		ret[i] = v * factor
	}
	return ret
}

// Divide ...
func Divide(a []float64, divider float64) []float64 {
	ret := make([]float64, len(a))
	copy(ret, a)

	for i, v := range ret {
		ret[i] = v / divider
	}
	return ret
}

// Dot ...
func Dot(a, b []float64) float64 {
	n1 := len(a)
	n2 := len(b)
	var n int
	switch n1 > n2 {
	case true:
		n = n2
	default:
		n = n1
	}
	ret := float64(0)

	for i := 0; i < n; i++ {
		ret += a[i] * b[i]
	}
	return ret
}