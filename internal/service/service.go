package service

import "github.com/lavatee/shop_api_gateway/internal/repository"

//go:generate mockgen -source=service.go -destination=mocks/mock.go -package=service

type Service struct {
	Auth
}

type Auth interface {
	SignUp(name string, email string, password string) (int, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo),
	}
}
