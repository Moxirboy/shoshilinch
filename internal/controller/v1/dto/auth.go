package dto



type SignUp struct{
	Firstname string `json:"first_name"`
	Lastname string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	ClassName string `json:"class_name"`
	Password string  `json:"password"`
}

type Login struct{
	PhoneNumber string `json:"phone_number"`
	Password string `json:"password"`
}
type AuthResponse struct{
	Token string `json:"token"`
	Role string  `json:"role"`
}
