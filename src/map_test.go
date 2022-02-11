package src

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonStr := `{"id":"P001","name":"Macbook Pro Max","image_url":"https://example.com/image.png"}`
	jsonBytes := []byte(jsonStr)

	product := map[string]interface{}{}

	err := json.Unmarshal(jsonBytes, &product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product)
	fmt.Println(product["id"])
}

func TestMapEncode(t *testing.T) {
	var product map[string]interface{} = map[string]interface{}{
		"id":    "P0001",
		"name":  "Iphone 12 Pro Max",
		"price": 10000000,
	}

	productByte, err := json.Marshal(product)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(productByte))
}
