package types

type User struct {
	Email    *string `json:"email" validate:"required,email"`
	Password *string `json:"password" validate:"required,gte=8,lte=50"`
}
