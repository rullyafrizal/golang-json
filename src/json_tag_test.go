package src

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Macbook Pro Max",
		ImageUrl: "https://example.com/image.png",
	}

	bytes, err := json.Marshal(product)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	var product Product = Product{}
	var jsonStr string = `{"id":"P001","name":"Macbook Pro Max","image_url":"https://example.com/image.png"}`
	var jsonBytes []byte = []byte(jsonStr)

	err := json.Unmarshal(jsonBytes, &product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
