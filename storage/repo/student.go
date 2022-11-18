package repo

import (
	"github.com/mirasildev/student/api/models"
	"time"
)

type Student struct {
	ID          int64
	FirstName   string
	LastName    string
	UserName    string
	Email       string
	PhoneNumber string
	CreatedAt   time.Time
}

// type InsertStudents struct {
// 	Students []*Student
// }

type GetAllStudentsParams struct {
	Limit  int32
	Page   int32
	Search string
}

type GetAllStudentsResult struct {
	Students []*Student
	Count int32
}

type StudentStorageI interface {
	Create(s *models.CreateStudent) (error)
	// GetAll(params *GetAllStudentsParams) (*GetAllStudentsResult, error)
}