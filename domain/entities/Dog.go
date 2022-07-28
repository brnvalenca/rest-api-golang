package entities

type Dog struct {
	Name         string `json:"name"`
	Age          string `json:"age"`
	MaleOrFemale string `json:"sex"`
	Color        string `json:"color"`
	Breed        Breed 
}
