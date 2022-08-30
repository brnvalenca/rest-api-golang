package entities

type DogBreed struct {
	BreedID      int    `json:"breed_id"`
	GoodWithKids int    `json:"goodwithkids"`
	GoodWithDogs int    `json:"goodwithdogs"`
	Shedding     int    `json:"shedding"`
	Grooming     int    `json:"grooming"`
	Energy       int    `json:"energy"`
	KennelID     int    `json:"kennelId"`
	Name         string `json:"name"`
	BreedImg     string `json:"imageurl"`
}
