package handler

import "shortener"

type signInInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LinkCreateRequest struct {
	URL string `json:"url" validate:"required,url"`
}

type getAllLinksResponse struct {
	Links []shortener.Link `json:"links"`
	Count int64            `json:"count"`
}
