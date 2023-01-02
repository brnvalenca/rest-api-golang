package entities

type User struct {
	ID              int    `json:"id"`
	Name            string `json:"name" validate:"required,min=2,max=128"`
	Email           string `json:"email" validate:"required, email"`
	Password        string `json:"password,omitempty" binding:"required"`
	UserPreferences UserDogPreferences
}

type UserBuilder struct {
	user *User
}

type UserAttrBuilder struct {
	UserBuilder
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{user: &User{}}
}

func (ub *UserBuilder) Has() *UserAttrBuilder {
	return &UserAttrBuilder{*ub}
}

func (ub *UserAttrBuilder) ID(id int) *UserAttrBuilder {
	ub.user.ID = id
	return ub
}

func (ub *UserAttrBuilder) Name(name string) *UserAttrBuilder {
	ub.user.Name = name
	return ub
}

func (ub *UserAttrBuilder) Email(email string) *UserAttrBuilder {
	ub.user.Email = email
	return ub
}

func (ub *UserAttrBuilder) Password(password string) *UserAttrBuilder {
	ub.user.Password = password
	return ub
}

func (ub *UserAttrBuilder) Uprefs(uprefs UserDogPreferences) *UserAttrBuilder {
	ub.user.UserPreferences = uprefs
	return ub
}

func (userbuilder *UserBuilder) BuildUser() *User {
	return userbuilder.user
}
