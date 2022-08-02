package entities

type DogKennel struct {
	Dogs    []Dog
	Address Address
	CNPJ    string `json:"cnpj"`
	Name    string `json:"name"`
}

func BuildDogKennel(d Dog, a Address, cnpj, name string) DogKennel {

	dogKennel := DogKennel{
		Dogs: []Dog{d},
		Address: Address{
			a.Street,
			a.District,
			a.PostalCode,
			a.City,
		},
		CNPJ: cnpj,
		Name: name,
	}

	return dogKennel
}
