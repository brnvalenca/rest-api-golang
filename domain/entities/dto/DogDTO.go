package dto

type DogDTO struct {

	/* Dog Info */
	KennelID  int    `json:"kennel_id"`
	BreedIDFK int    `json:"breed_idFK"`
	DogID     int    `json:"dog_id"`
	DogName   string `json:"name"`
	Sex       string `json:"sex"`
	/* breeds info */
	BreedID      int    `json:"breed_id"`
	GoodWithKids int    `json:"goodwithkids"`
	GoodWithDogs int    `json:"goodwithdogs"`
	Shedding     int    `json:"shedding"`
	Grooming     int    `json:"grooming"`
	Energy       int    `json:"energy"`
	BreedImg     string `json:"imageurl"`
}

func BuildDogDTO(dogname string, sex string, breedimg string, kennelid, breedid, dogid, gwithkds, gwithdgs, shed, groom, energy int) *DogDTO {
	dogDTO := DogDTO{
		DogName:      dogname,
		Sex:          sex,
		BreedImg:     breedimg,
		KennelID:     kennelid,
		BreedIDFK:    breedid,
		BreedID:      breedid,
		DogID:        dogid,
		GoodWithKids: gwithkds,
		GoodWithDogs: gwithdgs,
		Shedding:     shed,
		Grooming:     groom,
		Energy:       energy,
	}

	return &dogDTO
}
