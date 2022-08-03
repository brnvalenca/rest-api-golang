package entities

type DogKennel struct {
	Dogs    []Dog
	Address Address
	Name    string `json:"name"`
	ID      string    `json:"id"`
}

func BuildDogKennel(a Address, id, name string) DogKennel {
	

	dogKennel := DogKennel{
		ID:   id,
		Name: name,
		Address: Address{
			a.Street,
			a.District,
			a.PostalCode,
			a.City,
		},

	}

	return dogKennel
}

func (dk *DogKennel) AppendDogToKennel(d []Dog) []Dog {
	dk.Dogs = append(dk.Dogs, d...)
	return dk.Dogs
}
