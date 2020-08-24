package vectors

import "testing"

func TestAdd(t *testing.T) {
	ret := Add([]float64{1, 2, 3}, []float64{1, 2, 3})

	if ret[0] != 2 || ret[1] != 4 || ret[2] != 6 {
		t.Fatal("error")
	}
}

func TestSub(t *testing.T) {
	ret := Sub([]float64{1, 2, 3}, []float64{1, 2, 3})

	if ret[0] != 0 || ret[1] != 0 || ret[2] != 0 {
		t.Fatal("error")
	}
}

func TestDivide(t *testing.T) {
	ret := Divide([]float64{1, 2, 3}, 1)

	if ret[0] != 1 || ret[1] != 2 || ret[2] != 3 {
		t.Fatal("error")
	}
}

func TestMultiply(t *testing.T) {
	ret := Multiply([]float64{1, 2, 3}, 1)

	if ret[0] != 1 || ret[1] != 2 || ret[2] != 3 {
		t.Fatal("error")
	}
}

func TestDot(t *testing.T) {
	ret := Dot([]float64{1, 2, 3}, []float64{1, 2, 3})

	if ret != 14 {
		t.Fatal("error")
	}
}