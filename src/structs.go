package src

type Customer struct {
	Name      string
	Age       int
	Gender    string
	Hobbies   []string
	Addresses []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}
