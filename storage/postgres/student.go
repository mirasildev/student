package postgres

import (
	"github.com/mirasildev/student/api/models"

	"github.com/jmoiron/sqlx"
	"github.com/mirasildev/student/storage/repo"
)

type studentRepo struct {
	db *sqlx.DB
}

func NewStudent(db *sqlx.DB) repo.StudentStorageI {
	return &studentRepo{
		db: db,
	}
}

func (st *studentRepo) Create(student *models.CreateStudent) error {
	
	queryInsertStudent := `
		INSERT INTO student(
			first_name,
			last_name,
			username,
			email,
			phone_number
		) VALUES($1, $2, $3, $4, $5)
	`

	for _, students := range student.Students {
		_, err := st.db.Exec(
			queryInsertStudent,
			students.FirstName,
			students.LastName,
			students.UserName,
			students.Email,
			students.PhoneNumber,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// func (st *studentRepo) GetAllStudents()



