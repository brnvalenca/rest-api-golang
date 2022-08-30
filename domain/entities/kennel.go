package entities

type Kennel struct {
	ID     int        `json:"id"`
	Name   string     `json:"name"`
	Breeds []DogBreed `json:"breeds"`
}
