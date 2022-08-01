package entities

type Dog struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Color  string `json:"color"`
	IsMale bool   `json:"sex"`
	Age    int    `json:"age"`
	Breed  Breed
}

func BuildDog(b Breed, age int, ismale bool, name, color string, id string) Dog {

	d := Dog{

		Name:   name,
		Age:    age,
		IsMale: ismale,
		Color:  color,
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
