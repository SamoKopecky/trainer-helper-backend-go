package utils

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(i any) {
	b, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(b[:]))
}
