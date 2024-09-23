package request

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"email"`
}

type QueryParams struct {
	Email string `form:"email" binding:"omitempty,email"`
}
