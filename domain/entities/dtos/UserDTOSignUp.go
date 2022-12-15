package dtos

type UserDTOSignUp struct {
	ID        int    `json:"id"`
	Name      string `json:"name" validate:"required,min=2,max=128"`
	Email     string `json:"email" validate:"required, email"`
	Password  string `json:"-"`
	UserPrefs UserPrefsDTO
}

type UseDTOSignUpBuilder struct {
	user *UserDTOSignUp
}

type UserDTOSignUpAttrBuilder struct {
	UseDTOSignUpBuilder
}

func NewUserDTOBuilder() *UseDTOSignUpBuilder {
	return &UseDTOSignUpBuilder{user: &UserDTOSignUp{}}
}

func (ub *UseDTOSignUpBuilder) Has() *UserDTOSignUpAttrBuilder {
	return &UserDTOSignUpAttrBuilder{*ub}
}

func (ub *UserDTOSignUpAttrBuilder) ID(id int) *UserDTOSignUpAttrBuilder {
	ub.user.ID = id
	return ub
}

func (ub *UserDTOSignUpAttrBuilder) Name(name string) *UserDTOSignUpAttrBuilder {
	ub.user.Name = name
	return ub
}

func (ub *UserDTOSignUpAttrBuilder) Email(email string) *UserDTOSignUpAttrBuilder {
	ub.user.Email = email
	return ub
}

func (ub *UserDTOSignUpAttrBuilder) Password(password string) *UserDTOSignUpAttrBuilder {
	ub.user.Password = password
	return ub
}

func (ub *UserDTOSignUpAttrBuilder) UserPrefs(uprefs UserPrefsDTO) *UserDTOSignUpAttrBuilder {
	ub.user.UserPrefs = uprefs
	return ub
}

func (userbuilder *UseDTOSignUpBuilder) BuildUser() *UserDTOSignUp {
	return userbuilder.user
}
