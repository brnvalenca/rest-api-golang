package entities

type DogBreed struct {
	ID           int    `json:"breed_id"`
	GoodWithKids int    `json:"goodwithkids"`
	GoodWithDogs int    `json:"goodwithdogs"`
	Shedding     int    `json:"shedding"`
	Grooming     int    `json:"grooming"`
	Energy       int    `json:"energy"`
	Name         string `json:"name"`
	BreedImg     string `json:"imageurl"`
}

func BuildDogBreed(breedimg string, name string, breedid, dogid, gwithkds, gwithdgs, shed, groom, energy int) *DogBreed {
	dbreed := DogBreed{
		ID:           breedid,
		GoodWithDogs: gwithdgs,
		GoodWithKids: gwithkds,
		Shedding:     shed,
		Grooming:     groom,
		Energy:       energy,
		Name:         name,
		BreedImg:     breedimg,
	}

	return &dbreed
}
