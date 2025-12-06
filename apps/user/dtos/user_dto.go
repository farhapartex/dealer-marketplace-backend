package dtos

type CreateUserPayload struct {
	Name     string  `json:"name" binding:"required"`
	Email    *string `json:"email" binding:"omitempty,email"`
	Phone    *string `json:"phone"`
	Password string  `json:"password" binding:"required,min=6"`
}

type VerifyCodeRequest struct {
	Token string `json:"token" binding:"required"`
	Code  string `json:"code" binding:"required,len=6"`
}

type CreateUserResponse struct {
	Code  string `json:"code"`
	Token string `json:"token"`
}

type SigninRequest struct {
	Email    *string `json:"email" binding:"omitempty,email"`
	Mobile   *string `json:"mobile"`
	Password string  `json:"password" binding:"required"`
}

type SigninResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
