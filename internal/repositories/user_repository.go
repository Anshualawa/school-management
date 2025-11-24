package repositories

import (
	"github.com/Anshualawa/school-management/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct{ db *gorm.DB }

func NewUserRepository(db *gorm.DB) *UserRepository { return &UserRepository{db: db} }

func (r *UserRepository) Create(u *models.User) error { return r.db.Create(u).Error }

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var u models.User

	if err := r.db.Where("email=?", email).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) GetByID(id string) (*models.User, error) {
	var u models.User

	if err := r.db.Where("id=?", id).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}
