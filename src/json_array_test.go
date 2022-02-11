package src

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	cust := Customer{
		Name:    "Rully",
		Age:     19,
		Gender:  "Man",
		Hobbies: []string{"Swimming", "Diving", "Coding"},
	}

	bytes, _ := json.Marshal(cust)

	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"Name":"Rully","Age":19,"Gender":"Man","Hobbies":["Swimming","Diving","Coding"]}`
	jsonBytes := []byte(jsonString)

	cust := Customer{}

	err := json.Unmarshal(jsonBytes, &cust)

	if err != nil {
		panic(err)
	}

	fmt.Println(cust)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		Name: "Rully",
		Addresses: []Address{
			{Street: "Jl. Bougenvile no. 55", Country: "Indonesia", PostalCode: "61253"},
			{Street: "Jl. Bougenvile no. 54", Country: "Indonesia", PostalCode: "61252"},
		},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonStr := `{"Name":"Rully","Age":0,"Gender":"","Hobbies":null,"Addresses":[{"Street":"Jl. Bougenvile no. 55","Country":"Indonesia","PostalCode":"61253"},{"Street":"Jl. Bougenvile no. 54","Country":"Indonesia","PostalCode":"61252"}]}`
	jsonBytes := []byte(jsonStr)

	cust := Customer{}

	err := json.Unmarshal(jsonBytes, &cust)

	if err != nil {
		panic(err)
	}

	for _, v := range cust.Addresses {
		fmt.Println(v.PostalCode)
	}
}
