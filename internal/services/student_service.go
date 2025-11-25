package services

import (
	"github.com/Anshualawa/school-management/internal/models"
	"github.com/Anshualawa/school-management/internal/repositories"
	"github.com/google/uuid"
)

type StudentService struct {
	repo *repositories.StudentRepository
}

func NewStudentService(r *repositories.StudentRepository) *StudentService {
	return &StudentService{repo: r}
}

func (s *StudentService) CreateStudent(student *models.Student) (*models.Student, error) {
	student.ID = uuid.New()

	err := s.repo.Create(student)
	if err != nil {
		return nil, err
	}

	return student, nil

}

func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
	return s.repo.BetByID(id)
}
