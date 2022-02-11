package src

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonObject(t *testing.T) {
	cust := Customer{
		Name:   "Rully",
		Age:    19,
		Gender: "Man",
	}

	byted, err := json.Marshal(cust)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(byted))
}
