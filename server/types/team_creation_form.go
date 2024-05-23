package types

type Team struct {
	ID          *int64    `json:"id"` // ID is optional because it is not required when creating a new team
	Name        *string   `json:"name" validate:"required,min=2,max=100"`
	Email       *string   `json:"email" validate:"required,email"`
	Description *string   `json:"description" validate:"required"`
	Category    *int32    `json:"category" validate:"required"`
	City        *string   `json:"city" validate:"required"`
	Phone       *string   `json:"phone" validate:"required"`
	Gender      *string   `json:"gender" validate:"required"`
	Players     []*Player `json:"players" validate:"required,dive,required"`
}

type Player struct {
	FirstName *string `json:"first_name" validate:"required,min=2,max=100"`
	LastName  *string `json:"last_name" validate:"required,min=2,max=100"`
	Age       *int    `json:"age" validate:"required,gte=15,lte=100"`  // Assuming players are between 15 and 100 years old
	Number    *int    `json:"number" validate:"required,min=1,max=99"` // Assuming player numbers range from 1 to 99
	Gender    *string `json:"gender" validate:"required"`
	Position  *string `json:"position" validate:"required"`
	Height    *int    `json:"height" validate:"required,min=140,max=230"`   // Assuming player height (cm) range from 140 to 230
	Weight    *int    `json:"weight" validate:"required,min=30,max=200"`    // Assuming player weight (kg) range from 30 to 200
	Wingspan  *int    `json:"wingspan" validate:"required,min=100,max=250"` // Assuming player wingspan (cm) range from 100 to 250
	TeamID    uint64  `json:"teams_id"`
}
