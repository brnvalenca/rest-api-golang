package entities

type Dog struct {
	KennelID int      `json:"kennel_id"`
	BreedID  int      `json:"breed_id"`
	DogID    int      `json:"dog_id"`
	DogName  string   `json:"name"`
	Sex      string   `json:"sex"`
	Breed    DogBreed `json:"breed"`
}

type DogBuilder struct {
	dog *Dog
}

type DogAttrBuilder struct {
	DogBuilder
}

func NewDogBuilder() *DogBuilder {
	return &DogBuilder{dog: &Dog{}}
}

func (db *DogBuilder) Has() *DogAttrBuilder {
	return &DogAttrBuilder{*db}
}

func (db *DogAttrBuilder) KennelID(kennelID int) *DogAttrBuilder {
	db.dog.KennelID = kennelID
	return db
}

func (db *DogAttrBuilder) BreedID(breedID int) *DogAttrBuilder {
	db.dog.BreedID = breedID
	return db
}

func (db *DogAttrBuilder) DogID(dogID int) *DogAttrBuilder {
	db.dog.DogID = dogID
	return db
}

func (db *DogAttrBuilder) NameAndSex(dogname, dogsex string) *DogAttrBuilder {
	db.dog.DogName = dogname
	db.dog.Sex = dogsex
	return db
}

func (db *DogAttrBuilder) Breed(breed DogBreed) *DogAttrBuilder {
	db.dog.Breed = breed
	return db
}


func (dogbuilder *DogBuilder) BuildDog() *Dog {
	return dogbuilder.dog
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
		DogName:  name,
		Sex:      sex,
	}

	return &d
}
