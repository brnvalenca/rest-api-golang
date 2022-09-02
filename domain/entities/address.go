package entities

type Address struct {
	ID_Kennel int    `json:"kennel_id"`
	Numero    string `json:"numero"`
	Rua       string `json:"rua"`
	Bairro    string `json:"bairro"`
	CEP       string `json:"cep"`
	Cidade    string `json:"cidade"`
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
