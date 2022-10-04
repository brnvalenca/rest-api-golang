package entities

type User struct {
	ID              int    `json:"id"`
	Name            string `json:"name" validate:"required,min=2,max=128"`
	Email           string `json:"email" validate:"required, email"`
	Password        string `json:"password" validate:"passwd"`
	UserPreferences UserDogPreferences
}

/*
	Refactor:
	Nao trabalhar a entities na controller. Apenas na service e repository
	A controller não deve conhecer as entities, apenas as DTOs.
*/



func BuildUser(prefs UserDogPreferences, id int, name, email, password string) *User {

	u := User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
		UserPreferences: UserDogPreferences{
			prefs.UserID,
			prefs.GoodWithKids,
			prefs.GoodWithDogs,
			prefs.Shedding,
			prefs.Grooming,
			prefs.Energy,
		},
	}

	return &u
}
