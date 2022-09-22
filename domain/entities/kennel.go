package entities

type Kennel struct {
	ID            int    `json:"id"`
	ContactNumber string `json:"contact_number"`
	Name          string `json:"name"`
	Dogs          []Dog  `json:"dogs"`
	Address       Address
}

func BuildKennel(id int, dogs []Dog, addr Address, contact, name string) *Kennel {
	kennel := Kennel{
		ID:            id,
		ContactNumber: contact,
		Name:          name,
		Dogs:          dogs,
		Address: Address{
			addr.ID_Kennel,
			addr.Bairro,
			addr.CEP,
			addr.Cidade,
			addr.Numero,
			addr.Rua,
		},
	}
	return &kennel
}
