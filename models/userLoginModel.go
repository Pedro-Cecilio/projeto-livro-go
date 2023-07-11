package models

type UserLogin struct{
	Email string `json:"email" binding:"required,email"`
	Password CustomString `json:"password" binding:"required"`
}