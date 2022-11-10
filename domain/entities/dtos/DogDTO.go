package dtos

type DogDTO struct {
	KennelID int    `json:"kennel_id"`
	DogID    int    `json:"dog_id"`
	DogName  string `json:"name"`
	Sex      string `json:"sex"`

	BreedID      int    `json:"breed_id"`
	GoodWithKids int    `json:"goodwithkids"`
	GoodWithDogs int    `json:"goodwithdogs"`
	Shedding     int    `json:"shedding"`
	Grooming     int    `json:"grooming"`
	Energy       int    `json:"energy"`
	BreedName    string `json:"breed_name"`
	BreedImg     string `json:"imageurl"`
}

type DogDTOBuilder struct {
	dogDTO *DogDTO
}

type DogDTOAttrBuilder struct {
	DogDTOBuilder
}

func NewDogDTOBuilder() *DogDTOBuilder {
	return &DogDTOBuilder{dogDTO: &DogDTO{}}
}

func (db *DogDTOBuilder) Has() *DogDTOAttrBuilder {
	return &DogDTOAttrBuilder{*db}
}

func (db *DogDTOAttrBuilder) KennelID(kennelID int) *DogDTOAttrBuilder {
	db.dogDTO.KennelID = kennelID
	return db
}

func (db *DogDTOAttrBuilder) BreedID(breedID int) *DogDTOAttrBuilder {
	db.dogDTO.BreedID = breedID
	return db
}

func (db *DogDTOAttrBuilder) DogID(dogID int) *DogDTOAttrBuilder {
	db.dogDTO.DogID = dogID
	return db
}

func (db *DogDTOAttrBuilder) NameAndSex(dogname, dogsex string) *DogDTOAttrBuilder {
	db.dogDTO.DogName = dogname
	db.dogDTO.Sex = dogsex
	return db
}

func (db *DogDTOAttrBuilder) DogDTOBreedName(name string) *DogDTOAttrBuilder {
	db.dogDTO.BreedName = name
	return db
}

func (db *DogDTOAttrBuilder) DogDTOBreedImg(imgUrl string) *DogDTOAttrBuilder {
	db.dogDTO.BreedImg = imgUrl
	return db
}

func (db *DogDTOAttrBuilder) DogDTOGoodWithKidsAndDogs(gwithkids, gwithdogs int) *DogDTOAttrBuilder {
	db.dogDTO.GoodWithKids = gwithkids
	db.dogDTO.GoodWithDogs = gwithdogs
	return db
}

func (db *DogDTOAttrBuilder) DogDTOSheddGroomAndEnergy(shed, groom, energy int) *DogDTOAttrBuilder {
	db.dogDTO.Shedding = shed
	db.dogDTO.Grooming = groom
	db.dogDTO.Energy = energy
	return db
}

func (db *DogDTOBuilder) BuildDogDTO() *DogDTO {
	return db.dogDTO
}
