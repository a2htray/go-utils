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

	fmt.Println(StringMap(ss, func(s string) string {
		return s + "@"
	}))
}

func TestStringForeach(t *testing.T) {
	ss := []string{"1", "2", "3"}
	StringForeach(ss, func(s string) {
		fmt.Println(s)
	})
}

func TestStringUnique(t *testing.T) {
	ss := []string{"1", "2", "2"}
	fmt.Println(StringUnique(ss))
}

func TestStringExtend(t *testing.T) {
	ss := []string{"1", "2", "2"}
	s1 := []string{"4"}
	s2 := []string{"5"}
	fmt.Println(StringExtend(ss, s1, s2))
}