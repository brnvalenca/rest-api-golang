package entities

type Kennel struct {
	ID            int    `json:"id"`
	ContactNumber string `json:"contact_number"`
	Name          string `json:"name"`
	Dogs          []Dog  `json:"dogs"`
	Address       Address
}

type KennelBuilder struct {
	kennel *Kennel
}

type KennelAttrBuilder struct {
	KennelBuilder
}

func NewKennelBuilder() *KennelBuilder {
	return &KennelBuilder{kennel: &Kennel{}}
}

func (kb *KennelBuilder) Has() *KennelAttrBuilder {
	return &KennelAttrBuilder{*kb}
}

func (kb *KennelAttrBuilder) ID(id int) *KennelAttrBuilder {
	kb.kennel.ID = id
	return kb
}

func (kb *KennelAttrBuilder) ContactNumber(contactNumber string) *KennelAttrBuilder {
	kb.kennel.ContactNumber = contactNumber
	return kb
}

func (kb *KennelAttrBuilder) Name(name string) *KennelAttrBuilder {
	kb.kennel.Name = name
	return kb
}

func (kb *KennelAttrBuilder) Dogs(dogs []Dog) *KennelAttrBuilder {
	kb.kennel.Dogs = dogs
	return kb
}

func (kb *KennelAttrBuilder) Address(addr Address) *KennelAttrBuilder {
	kb.kennel.Address = addr
	return kb
}

func (kb *KennelBuilder) BuildKennel() *Kennel {
	return kb.kennel
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
