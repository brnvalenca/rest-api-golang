package entities

type UserDogPreferences struct {
	DogLoudness int    `json:"noise"`
	DogEnergy   int    `json:"energy"`
	DogAVGSize   string `json:"age"`

}

func BuildUserDogPreferences(loudness int, energy int, avgsize string) UserDogPreferences {

	udog := UserDogPreferences{
		DogLoudness: loudness,
		DogEnergy:   energy,
		DogAVGSize:   avgsize,

	}

	return udog
}
