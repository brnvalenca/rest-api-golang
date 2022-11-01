package dtos

type BreedDTO struct {
	ID           int    `json:"breed_id"`
	GoodWithKids int    `json:"goodwithkids"`
	GoodWithDogs int    `json:"goodwithdogs"`
	Shedding     int    `json:"shedding"`
	Grooming     int    `json:"grooming"`
	Energy       int    `json:"energy"`
	Name         string `json:"name"`
	BreedImg     string `json:"imageurl"`
}

// Fazer um construtor pra esse DTO

type DogBreedBuilder struct {
	dogbreed *BreedDTO
}

type BreedAttrBuilder struct {
	DogBreedBuilder
}

func NewBreedBuilderDTO() *DogBreedBuilder {
	return &DogBreedBuilder{dogbreed: &BreedDTO{}}
}

func (d *DogBreedBuilder) Has() *BreedAttrBuilder {
	return &BreedAttrBuilder{*d}
}

func (attr *BreedAttrBuilder) ID(id int) *BreedAttrBuilder {
	attr.dogbreed.ID = id
	return attr
}

func (attr *BreedAttrBuilder) Name(name string) *BreedAttrBuilder {
	attr.dogbreed.Name = name
	return attr
}

func (attr *BreedAttrBuilder) Img(imgUrl string) *BreedAttrBuilder {
	attr.dogbreed.BreedImg = imgUrl
	return attr
}

func (attr *BreedAttrBuilder) GoodWithKidsAndDogs(gwithkids, gwithdogs int) *BreedAttrBuilder {
	attr.dogbreed.GoodWithKids = gwithkids
	attr.dogbreed.GoodWithDogs = gwithdogs
	return attr
}

func (attr *BreedAttrBuilder) SheddGroomAndEnergy(shed, groom, energy int) *BreedAttrBuilder {
	attr.dogbreed.Shedding = shed
	attr.dogbreed.Grooming = groom
	attr.dogbreed.Energy = energy
	return attr
}

func (db *DogBreedBuilder) BuildBreedDTO() *BreedDTO {
	return db.dogbreed
}

func BuildDogBreedDTO(breedimg string, name string, breedid, dogid, gwithkds, gwithdgs, shed, groom, energy int) *BreedDTO {
	dbreed := BreedDTO{
		ID:           breedid,
		Name:         name,
		GoodWithDogs: gwithdgs,
		GoodWithKids: gwithkds,
		Shedding:     shed,
		Grooming:     groom,
		Energy:       energy,
		BreedImg:     breedimg,
	}

	return &dbreed
}
