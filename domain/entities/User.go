package entities

var UsersData = make([]User, 0)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func BuildUser(name string, email string, password string) User {

	u := User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	return u
	
}