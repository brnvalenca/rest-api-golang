package dtos

type UserDTOSignIn struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"passwd"`
}
