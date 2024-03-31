package dto


type CreateTeacher struct{
	Firstname string `json:"first_name" binding:"required"`
	Lastname string `json:"last_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Teacher struct{
	ID string `json:"id"`
	Firstname string `json:"first_name"`
	Lastname string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Password string `json:"password"`
	Role string `json:"role"`

}