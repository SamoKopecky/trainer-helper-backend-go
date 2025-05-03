package utils

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(i any) {
	b, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(b[:]))
}

func PrettyString(i any) string {
	b, _ := json.MarshalIndent(i, "", "\t")
	return fmt.Sprint(string(b[:]))
}
