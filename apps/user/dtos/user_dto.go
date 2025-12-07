package dtos

type CreateUserPayload struct {
	Name     string  `json:"name" binding:"required"`
	Email    *string `json:"email" binding:"omitempty,email"`
	Phone    *string `json:"phone"`
	Password string  `json:"password" binding:"required,min=6"`
	UserType string  `json:"user_type" binding:"required,oneof=dealer sub_dealer customer"`
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

type UserMeResponse struct {
	UserID            string    `json:"user_id"`
	Name              string    `json:"name"`
	Email             *string   `json:"email"`
	Phone             *string   `json:"phone"`
	UserType          *string   `json:"user_type"`
	IsOnboardComplete bool      `json:"is_onboard_complete"`
	Shop              *ShopInfo `json:"shop,omitempty"`
}

type ShopInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
