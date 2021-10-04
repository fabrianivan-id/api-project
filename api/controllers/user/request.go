package user

type PostUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	//Gender   string `json:"gender" form:"gender"`
	Password string `json:"password" form:"password"`
}

type EditUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	//Gender   string `json:"gender" form:"gender"`
	Password string `json:"password" form:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
