package entities

type Dog struct {
	KennelID int      `json:"kennel_id"`
	BreedID  int      `json:"breed_id"`
	DogID    int      `json:"dog_id"`
	DogName     string   `json:"name"`
	Sex      string   `json:"sex"`
	Breed    DogBreed `json:"breed"`
}

func BuildDog(breed DogBreed, dogid int, kennelid int, sex, name string) *Dog {

	d := Dog{
		Breed: DogBreed{
			breed.ID,
			breed.GoodWithKids,
			breed.GoodWithDogs,
			breed.Shedding,
			breed.Grooming,
			breed.Energy,
			breed.Name,
			breed.BreedImg,
		},
		KennelID: kennelid,
		DogID:    dogid,
		DogName:     name,
		Sex:      sex,
	}

	return &d
}
