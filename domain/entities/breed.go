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

type DogBreedBuilder struct {
	dogbreed *DogBreed
}

type BreedAttrBuilder struct {
	DogBreedBuilder
}

func NewDogBreedBuilder() *DogBreedBuilder {
	return &DogBreedBuilder{dogbreed: &DogBreed{}}
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

func (db *DogBreedBuilder) BuildBreed() *DogBreed {
	return db.dogbreed
}

func BuildDogBreed(breedimg string, name string, breedid, dogid, gwithkds, gwithdgs, shed, groom, energy int) *DogBreed {
	dbreed := DogBreed{
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
