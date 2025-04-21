package repositories

import "database/sql"

type Repository interface {
	GetMessage() (string, error)
}

type RepoImpl struct {
	db *sql.DB
}

func NewRepository() Repository {
	return &RepoImpl{}
}

func (r *RepoImpl) GetMessage() (string, error) {
	return " data from repository", nil
}
