package pretty

import (
	"encoding/json"
	"fmt"
)

func Print(outs ...interface{}) {
	output, err := json.MarshalIndent(outs, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}