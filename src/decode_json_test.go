package src

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonRequset := `{"Name":"Rully","Age":19,"Gender":"Man"}`

	jsonBytes := []byte(jsonRequset)

	customer := Customer{}

	err := json.Unmarshal(jsonBytes, &customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}