package slices

import (
	"fmt"
	"strconv"
	"testing"
)

func TestInterfaceForeach(t *testing.T) {
	ss := []interface{}{
		"1",
		"2",
		"3",
	}
	newSS := make([]int, 0, len(ss))
	InterfaceForeach(ss, func(i interface{}) {
		o, _ := strconv.Atoi(i.(string))
		newSS = append(newSS, o)
	})
	fmt.Println(newSS)
}
