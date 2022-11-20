package postgres

import (
	"github.com/mirasildev/student/api/models"

	"fmt"

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

func (st *studentRepo) GetAll(params *repo.GetAllStudentsParams) (*repo.GetAllStudentsResult, error) {
	result := repo.GetAllStudentsResult{
		Students: make([]*repo.Student, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	filter := ""
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter += fmt.Sprintf(`
			WHERE first_name ILIKE '%s' OR last_name ILIKE '%s' OR email ILIKE '%s' 
				OR username ILIKE '%s' OR phone_number ILIKE '%s'`,
			str, str, str, str, str,
		)
	}

	query := `
		SELECT
			id,
			first_name,
			last_name,
			username,
			email,
			phone_number,
			created_at
		FROM student
		` + filter + `
		ORDER BY created_at desc
		` + limit

	rows, err := st.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var s repo.Student

		err := rows.Scan(
			&s.ID,
			&s.FirstName,
			&s.LastName,
			&s.UserName,
			&s.Email,
			&s.PhoneNumber,
			&s.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		result.Students = append(result.Students, &s)
	}

	queryCount := `SELECT count(1) FROM student ` + filter
	err = st.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
