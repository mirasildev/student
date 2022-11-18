package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/mirasildev/student/storage/postgres"
	"github.com/mirasildev/student/storage/repo"
)

type StorageI interface {
	Student() repo.StudentStorageI
}

type storagePg struct {
	studentRepo repo.StudentStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		studentRepo: postgres.NewStudent(db),
	}
}

func (s *storagePg) Student() repo.StudentStorageI {
	return s.studentRepo
}
