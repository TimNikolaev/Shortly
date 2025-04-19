package handler

type signInInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LinkCreateRequest struct {
	URL string `json:"url" validate:"required,url"`
}
