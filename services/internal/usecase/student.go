package usecase

import (
	"go-web-robotek/services/internal/domain"
	"go-web-robotek/services/internal/repository"
)

type Student interface {
	Create(fullName, email, password, phoneNumber string, groupID int) (int, error)
	GetOne(id int) (*domain.Student, error)
	GetAll() (*[]domain.Student, error)
	GetGroup(id int) (*domain.Group, error)
	Delete(id int) (int, error)
	ChangeGroup(studentID, groupID int) error
	Update(id int, updatedStudent domain.Student) error
}

type StudentUseCase struct {
	studentRepo repository.StudentRepo
}

func NewStudentUsecase(studentRepo repository.StudentRepo) *StudentUseCase {
	return &StudentUseCase{
		studentRepo: studentRepo,
	}
}

func (uc *StudentUseCase) Create(fullName, email, password, phoneNumber string, groupID int) (int, error) {
	student := domain.Student{
		FullName:    fullName,
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
		GroupID:     groupID,
	}

	id, err := uc.studentRepo.Create(&student)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *StudentUseCase) GetOne(id int) (*domain.Student, error) {
	student, err := uc.studentRepo.GetOne(id)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (uc *StudentUseCase) GetGroup(id int) (*domain.Group, error) {
	group, err := uc.studentRepo.GetGroup(id)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (uc *StudentUseCase) GetAll() (*[]domain.Student, error) {
	students, err := uc.studentRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return &students, nil
}

func (uc *StudentUseCase) Delete(id int) (int, error) {
	deletedID, err := uc.studentRepo.Delete(id)
	if err != nil {
		return 0, err
	}

	return deletedID, nil
}

func (uc *StudentUseCase) ChangeGroup(studentID, groupID int) error {
	err := uc.studentRepo.ChangeGroup(studentID, groupID)
	if err != nil {
		return err
	}

	return nil
}

func (uc *StudentUseCase) Update(id int, updatedStudent domain.Student) error {
	err := uc.studentRepo.Update(id, &updatedStudent)
	if err != nil {
		return err
	}

	if updatedStudent.GroupID != 0 {
		err = uc.studentRepo.ChangeGroup(id, updatedStudent.GroupID)
		if err != nil {
			return err
		}
	}

	return nil
}