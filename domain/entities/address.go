package entities

type Address struct {
	ID_Kennel int    `json:"kennel_id"`
	Numero    string `json:"numero"`
	Rua       string `json:"rua"`
	Bairro    string `json:"bairro"`
	CEP       string `json:"cep"`
	Cidade    string `json:"cidade"`
}

type AddressBuilder struct {
	addr *Address
}

type AddressAttr struct {
	AddressBuilder
}

func NewAddressBuilder() *AddressBuilder {
	return &AddressBuilder{addr: &Address{}}
}

func (addrBuilder *AddressBuilder) Has() *AddressAttr {
	return &AddressAttr{*addrBuilder}
}

func (addrAttr *AddressAttr) IDKennel(idkennel int) *AddressAttr {
	addrAttr.addr.ID_Kennel = idkennel
	return addrAttr
}

func (addrAttr *AddressAttr) Numero(numero string) *AddressAttr {
	addrAttr.addr.Numero = numero
	return addrAttr
}

func (addrAttr *AddressAttr) Rua(rua string) *AddressAttr {
	addrAttr.addr.Rua = rua
	return addrAttr
}

func (addrAttr *AddressAttr) Bairro(bairro string) *AddressAttr {
	addrAttr.addr.Bairro = bairro
	return addrAttr
}

func (addrAttr *AddressAttr) CEP(cep string) *AddressAttr {
	addrAttr.addr.CEP = cep
	return addrAttr
}

func (addrAttr *AddressAttr) Cidade(cidade string) *AddressAttr {
	addrAttr.addr.Cidade = cidade
	return addrAttr
}

func (addr *AddressBuilder) BuildAddr() *Address {
	return addr.addr
}

func BuildAddress(idKennel int, num, rua, bairro, cep, city string) *Address {
	addr := Address{
		ID_Kennel: idKennel,
		Numero:    num,
		Rua:       rua,
		Bairro:    bairro,
		CEP:       cep,
		Cidade:    city,
	}

	return &addr
}
