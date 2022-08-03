package entities

type Dog struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Sex   string `json:"sex"`
	Age   int    `json:"age"`
	Breed Breed
}

func BuildDog(b Breed, age int, sex, name, id string) Dog {

	d := Dog{
		Name: name,
		Age:  age,
		Sex:  sex,
		Breed: Breed{
			b.BreedName,
			b.BreedAVGSize,
			b.BreedLoudness,
			b.BreedEnergy,
		},
		ID: id,
	}

	return d
}
