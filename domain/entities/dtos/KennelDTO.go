package dtos

type KennelDTO struct {
	ID            int    `json:"id"`
	ContactNumber string `json:"contact_number"`
	Name          string `json:"name"`

	Numero string `json:"numero"`
	Rua    string `json:"rua"`
	Bairro string `json:"bairro"`
	CEP    string `json:"cep"`
	Cidade string `json:"cidade"`

	Dogs []DogDTO `json:"dog_list"`
}

type KennelBuilder struct {
	kenneldto *KennelDTO
}

type KennelAttrBuilder struct {
	KennelBuilder
}

func NewKennelBuilderDTO() *KennelBuilder {
	return &KennelBuilder{kenneldto: &KennelDTO{}}
}

func (kb *KennelBuilder) Has() *KennelAttrBuilder {
	return &KennelAttrBuilder{*kb}
}

func (kb *KennelAttrBuilder) ID(id int) *KennelAttrBuilder {
	kb.kenneldto.ID = id
	return kb
}

func (kb *KennelAttrBuilder) ContactNumber(contactNumber string) *KennelAttrBuilder {
	kb.kenneldto.ContactNumber = contactNumber
	return kb
}

func (kb *KennelAttrBuilder) Name(name string) *KennelAttrBuilder {
	kb.kenneldto.Name = name
	return kb
}

func (kb *KennelAttrBuilder) Numero(numero string) *KennelAttrBuilder {
	kb.kenneldto.Numero = numero
	return kb
}

func (kb *KennelAttrBuilder) Rua(rua string) *KennelAttrBuilder {
	kb.kenneldto.Rua = rua
	return kb
}

func (kb *KennelAttrBuilder) Bairro(bairro string) *KennelAttrBuilder {
	kb.kenneldto.Bairro = bairro
	return kb
}

func (kb *KennelAttrBuilder) CEP(cep string) *KennelAttrBuilder {
	kb.kenneldto.CEP = cep
	return kb
}

func (kb *KennelAttrBuilder) Cidade(cidade string) *KennelAttrBuilder {
	kb.kenneldto.Cidade = cidade
	return kb
}

func (kb *KennelBuilder) BuildKennel() *KennelDTO {
	return kb.kenneldto
}
