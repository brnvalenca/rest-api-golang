package entities

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func BuildUser(id string, name string, email string, password string) User {

	u := User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}

	return u

}
