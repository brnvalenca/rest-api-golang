package entities

type Address struct {
	Street     string `json:"street"`
	District   string `json:"district"`
	PostalCode string `json:"postalcode"`
	City       string `json:"city"`
}

func BuildAddress(street, district, postalcode, city string) Address {
	a := Address{
		Street:     street,
		District:   district,
		PostalCode: postalcode,
		City:       city,
	}

	return a
}
