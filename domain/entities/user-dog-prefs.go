package entities

type UserDogPreferences struct {
	UserID       int `json:"user_id"`
	GoodWithKids int `json:"good_with_kids"`
	GoodWithDogs int `json:"good_with_dogs"`
	Shedding     int `json:"shedding"`
	Grooming     int `json:"grooming"`
	Energy       int `json:"energy"`
}

func BuildUserDogPreferences(id, gdwithkids, gdwithdogs, shedd, groom, energy int) UserDogPreferences {

	udogpref := UserDogPreferences{
		UserID:       id,
		GoodWithKids: gdwithkids,
		GoodWithDogs: gdwithdogs,
		Shedding:     shedd,
		Grooming:     groom,
		Energy:       energy,
	}
	return udogpref
}
