package usecase

import (
	"go-web-robotek/services/internal/domain"
	"go-web-robotek/services/internal/repository"
)

type Teacher interface {
	Create(fullName, email, password, phoneNumber string) (int, error)
	GetOne(id int) (*domain.Teacher, error)
	GetAll() (*[]domain.Teacher, error)
}

type TeacherUseCase struct {
	teacherRepo repository.TeacherRepo
}

func NewTeacherUsecase(teacherRepo repository.TeacherRepo) *TeacherUseCase {
	return &TeacherUseCase{
		teacherRepo: teacherRepo,
	}
}

func (uc *TeacherUseCase) Create(fullName, email, password, phoneNumber string) (int, error) {
	teacher := domain.Teacher{
		FullName:    fullName,
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
	}

	id, err := uc.teacherRepo.Create(&teacher)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *TeacherUseCase) GetOne(id int) (*domain.Teacher, error) {
	teacher, err := uc.teacherRepo.GetOne(id)
	if err != nil {
		return nil, err
	}

	return &teacher, nil
}

func (uc *TeacherUseCase) GetAll() (*[]domain.Teacher, error) {
	teachers, err := uc.teacherRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return &teachers, nil
}
