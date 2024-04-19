package types

type User struct {
	FirstName *string `json:"first_name" validate:"required,min=2,max=100"`
	LastName  *string `json:"last_name" validate:"omitempty,min=2,max=100"`
	Email     *string `json:"email" validate:"required,email"`
	Password  *string `json:"password" validate:"required,gte=8"`
}
