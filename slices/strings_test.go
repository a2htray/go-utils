package slices

import (
	"fmt"
	"testing"
)

func TestStrIn(t *testing.T) {
	fmt.Println(StringIn([]string{"1", "2", "3"}, "3"))
}

func TestStrMap(t *testing.T) {
	ss := []string{"1", "2", "3"}

	fmt.Println(StringMap(ss, func(s string) interface{} {
		return s + "@"
	}))
}

func TestStringForeach(t *testing.T) {
	ss := []string{"1", "2", "3"}
	StringForeach(ss, func(s string) {
		fmt.Println(s)
	})
}