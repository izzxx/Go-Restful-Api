package user

type UserResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type UserUpdatePassword struct {
	Email        string `json:"email"`
	PastPassword string `json:"past_password"`
	NewPassword  string `json:"new_password"`
}
