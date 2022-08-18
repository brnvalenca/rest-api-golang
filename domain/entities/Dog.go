package entities

type Dog struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Sex   string `json:"sex"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
}

func BuildDog(age int, breed, sex, name, id string) Dog {

	d := Dog{
		Name:  name,
		Age:   age,
		Sex:   sex,
		Breed: breed,
		ID:    id,
	}

	return d
}
