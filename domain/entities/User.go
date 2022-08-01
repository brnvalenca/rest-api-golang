package entities

type User struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"-"`
	Address        Address
	DogPreferences UserDogPreferences
}

func BuildUser(id string, name string, email string, password string, a Address, udog UserDogPreferences) User {

	u := User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
		Address: Address{
			a.Street,
			a.District,
			a.PostalCode,
			a.City,
		},
		DogPreferences: UserDogPreferences{
			udog.DogLoudness,
			udog.DogEnergy,
			udog.DogAVGSize,
		},
	}

	return u

}
