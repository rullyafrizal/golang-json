package src

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDecoder(t *testing.T) {
	reader, _ := os.Open("customer.json")
	decoder := json.NewDecoder(reader)

	customer := Customer{}

	err := decoder.Decode(&customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("product.json")
	encoder := json.NewEncoder(writer)

	product := Product{
		Id:       "P001",
		Name:     "Macbook Pro Max",
		ImageUrl: "https://example.com/image.png",
	}

	err := encoder.Encode(&product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
