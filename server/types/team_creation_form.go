package types

type Team struct {
	ID          *uint64   `json:"id"`
	CategoryID  *uint64   `json:"category_id" validate:"required"`
	CityID      *uint64   `json:"city_id" validate:"required"`
	Name        *string   `json:"name" validate:"required,min=2,max=100"`
	Description *string   `json:"description" validate:"required"`
	Email       *string   `json:"email" validate:"required,email"`
	Phone       *string   `json:"phone" validate:"required"`
	Players     *[]Player `json:"players" validate:"required"`
}

type Player struct {
	FirstName *string `json:"first_name" validate:"required,min=2,max=100"`
	LastName  *string `json:"last_name" validate:"required,min=2,max=100"`
	Age       *int    `json:"age" validate:"required,gte=1,lte=100"`
	Number    *int    `json:"number" validate:"required,min=0,max=99"`
	Gender    *string `json:"gender" validate:"required"`
	Position  *string `json:"position" validate:"required"`
	Height    *int    `json:"height" validate:"required,min=140,max=230"`
	Weight    *int    `json:"weight" validate:"required,min=30,max=200"`
	Wingspan  *int    `json:"wingspan" validate:"required,min=100,max=250"`
}
