package src

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(v interface{}) {
	byted, err := json.Marshal(v)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(byted))
}

func TestEncode(t *testing.T) {
	logJson("Rully")
	logJson(123)
	logJson(true)
	logJson([]string{"Rully", "Afrizal", "Alwin"})
}