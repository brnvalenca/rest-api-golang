package entities

type Breed struct {
	BreedName     string `json:"breename"`
	BreedAVGSize  string `json:"avgsize"`
	BreedLoudness string `json:"loudness"`
	BreedEnergy   string `json:"energy"`
}

func BuildBreed(breedname string, avgsize string, loudness string, energy string) Breed {

	b := Breed{
		BreedName:     breedname,
		BreedAVGSize:  avgsize,
		BreedLoudness: loudness,
		BreedEnergy:   energy,
	}

	return b
}
