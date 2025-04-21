
package services

import "fiber-postgres/internal/repositories"

type Service interface {
	GetMessage() (string, error)
}

type ServiceImpl struct {
	repo repositories.Repository
}

func NewService(r repositories.Repository) Service {
	return &ServiceImpl{repo: r}
}

func (s *ServiceImpl) GetMessage() (string, error) {
	return s.repo.GetMessage()
}

