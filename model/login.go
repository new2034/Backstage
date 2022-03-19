package model

type Login struct {
	Email string `json:"email" binding:"email"`
	Password string `json:"password"`
}