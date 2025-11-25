package repositories

import (
	"github.com/Anshualawa/school-management/internal/models"
	"gorm.io/gorm"
)

type StudentRepository struct{ db *gorm.DB }

func NewStudentRepository(db *gorm.DB) *StudentRepository { return &StudentRepository{db: db} }

func (r *StudentRepository) Create(s *models.Student) error { return r.db.Create(s).Error }

func (r *StudentRepository) GetByRollNo(role_number int) (*models.Student, error) {
	var s models.Student

	if err := r.db.Where("rollno=?", role_number).First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *StudentRepository) BetByID(id string) (*models.Student, error) {
	var s models.Student
	if err := r.db.Where("user_id=?", id).First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}
