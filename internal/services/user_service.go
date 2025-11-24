package services

import (
	"time"

	"github.com/Anshualawa/school-management/internal/auth"
	"github.com/Anshualawa/school-management/internal/models"
	"github.com/Anshualawa/school-management/internal/repositories"
	"github.com/Anshualawa/school-management/internal/utils"
	"github.com/google/uuid"
)

type UserService struct{ repo *repositories.UserRepository }

func NewUserService(r *repositories.UserRepository) *UserService { return &UserService{repo: r} }

func (s *UserService) Register(name, email, password, role string) (*models.User, error) {
	u := &models.User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	hashedPass, err := auth.HashedPassword(password)

	if err != nil {
		return nil, err
	}

	u.Password = hashedPass

	if utils.IsEmpty(role) {
		u.Role = "student"
	}

	if err := s.repo.Create(u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *UserService) Login(email, password string) (*models.User, string, error) {

	u, err := s.repo.GetByEmail(email)

	if err != nil {
		return nil, "", err
	}

	if !auth.CheckPassword(u.Password, password) {
		return nil, "", err
	}

	token, _ := auth.GenerateJWT(u.ID.String(), u.Name, u.Email, u.Role)

	return u, token, nil
}
