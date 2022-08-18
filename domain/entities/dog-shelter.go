package entities

type DogShelter struct {
	Dogs    []Dog
	Address Address
	Name    string `json:"name"`
	ID      string `json:"id"`
}

func BuildDogShelter(a Address, id, name string) DogShelter {

	dogShelter := DogShelter{
		ID:   id,
		Name: name,
		Address: Address{
			a.Street,
			a.District,
			a.PostalCode,
			a.City,
		},
	}

	return dogShelter
}

func (dk *DogShelter) AppendDogToShelter(d []Dog) []Dog {
	dk.Dogs = append(dk.Dogs, d...)
	return dk.Dogs
}
