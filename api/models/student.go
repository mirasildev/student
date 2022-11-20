package models

import "time"

type Student struct {
	ID          int64 `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}
type StudentRequest struct {
	FirstName   string `json:"first_name" binding:"required,min=2,max=30"`
	LastName    string `json:"last_name" binding:"required,min=2,max=30"`
	Email       string `json:"email" binding:"required,email"`
	UserName    string `json:"username"`
	PhoneNumber *string `json:"phone_number"`
}


type CreateStudent struct {
	Students []*StudentRequest `json:"student" binding:"required,dive"`
}

type GetAllStudentsResponse struct {
	Students []*Student `json:"students"`
	Count int32   `json:"count"`
}

type GetAllParams struct {
	Limit  int32  `json:"limit" binding:"required" default:"10"`
	Page   int32  `json:"page" binding:"required" default:"1"`
	Search string `json:"search"`
}