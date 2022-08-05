package entities

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func BuildUser(id int, name, email, password string) User {

	u := User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}

	return u

}
