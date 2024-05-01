package types

type Team struct {
	Name        *string   `json:"name" validate:"required,min=2,max=100"`
	Email       *string   `json:"email" validate:"required,email"`
	Description *string   `json:"description" validate:"required"`
	Category    *string   `json:"category" validate:"required"`
	City        *string   `json:"city" validate:"required"`
	Players     []*Player `json:"players" validate:"required,dive,required"`
}

type Player struct {
	FirstName *string `json:"first_name" validate:"required,min=2,max=100"`
	LastName  *string `json:"last_name" validate:"required,min=2,max=100"`
	Age       *int    `json:"age" validate:"required,gte=15,lte=100"`  // Assuming players are between 15 and 100 years old
	Number    *int    `json:"number" validate:"required,min=1,max=99"` // Assuming player numbers range from 1 to 99
}
