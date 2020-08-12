package strings

import (
	"fmt"
	"testing"
)

func TestReSplit(t *testing.T) {
	s := "a2htray@.#yuen"
	for _, v := range RegexSplit(s, `[@.#]+`, -1)  {
		fmt.Println(v)
	}
}

func main() {

}