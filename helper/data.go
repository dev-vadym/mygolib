package helper

import (
	"encoding/json"
	"fmt"
)

func Dump(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}
