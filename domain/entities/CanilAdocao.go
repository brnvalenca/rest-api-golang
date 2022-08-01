package entities

type CanilAdocao struct {
	CNPJ string `json:"cnpj"`
	Name string `json:"name"`
	Dogs []Dog  
}

