package pretty

import "testing"

func TestPrint(t *testing.T) {
	type V struct {
		A int `json:"a"`
		B string `json:"b"`
	}
	a := []interface{}{V{
		A: 1,
		B: "a",
	}}
	Print(a...)

	b := []interface{}{1, 2, 3}
	Print(b...)
}