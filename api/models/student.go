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
	Students []*StudentRequest
}

type GetAllUsersResponse struct {
	Students []*Student `json:"categories"`
	Count int32   `json:"count"`
}