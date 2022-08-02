package entities

type DogKennel struct {
	Dogs    []Dog
	Address Address
	CNPJ    string `json:"cnpj"`
	Name    string `json:"name"`
	ID      int    `json:"id"`
}

func BuildDogKennel(a Address, id int, cnpj, name string) DogKennel {

	dogKennel := DogKennel{
		Address: Address{
			a.Street,
			a.District,
			a.PostalCode,
			a.City,
		},
		ID:   id,
		CNPJ: cnpj,
		Name: name,
	}

	return dogKennel
}

func (dk *DogKennel) AppendDogToKennel(d []Dog) []Dog {
	dk.Dogs = append(dk.Dogs, d...)
	return dk.Dogs
}
