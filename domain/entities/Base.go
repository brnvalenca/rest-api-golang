package entities

import "time"

type Base struct {
	ID UniqueEntityID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

/* 
	Ele basicamente diz que essa clase vai ter os dados basicos que tem que ter em outras entidades
	ai a gente deixa tudo guardado aqui e as outras pegam esse tipo Base e usam. 
*/