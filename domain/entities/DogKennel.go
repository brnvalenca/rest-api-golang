package entities

type DogKennel struct {
	Dogs    []Dog
	Address Address
	CNPJ    string `json:"cnpj"`
	Name    string `json:"name"`
}
