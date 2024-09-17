package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/lavatee/shop_api_gateway/internal/repository"
)

const (
	salt = "x9q47sa5n"
)

type AuthService struct {
	Repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{
		Repo: repo,
	}
}

func (s *AuthService) SignUp(name string, email string, password string) (int, error) {
	hash := sha1.New()
	hash.Write([]byte(password))
	passwordHash := fmt.Sprintf("%s", hash.Sum([]byte(salt)))
	return s.Repo.SignUp(name, email, passwordHash)
}
