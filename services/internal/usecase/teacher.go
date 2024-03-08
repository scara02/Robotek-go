package usecase

import (
	"go-web-robotek/services/internal/domain"
	"go-web-robotek/services/internal/repository"
)

type Teacher interface {
	Create(fullName, email, password, phoneNumber string) (int, error)
	GetOne(id int) (*domain.Teacher, error)
	GetAll() (*[]domain.Teacher, error)
	Delete(id int) (int, error)
	AddToGroup(teacherID, groupID int) (error)
	GetGroups(id int) ([]domain.Group, error)
	Update(id int, updatedTeacher domain.Teacher) error
	DeleteGroup(teacherID, groupID int) error
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

func (uc *TeacherUseCase) Delete(id int) (int, error) {
	deletedID, err := uc.teacherRepo.Delete(id)
	if err != nil {
		return 0, err
	}

	return deletedID, nil
}

func (uc *TeacherUseCase) AddToGroup(teacherID, groupID int) (error) {
	err := uc.teacherRepo.AddToGroup(teacherID, groupID)
	if err != nil {
		return err
	}

	return nil
}

func (uc *TeacherUseCase) GetGroups(id int) ([]domain.Group, error) {
	groups, err := uc.teacherRepo.GetGroups(id)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (uc *TeacherUseCase) DeleteGroup(teacherID, groupID int) error {
	err := uc.teacherRepo.DeleteGroup(teacherID, groupID)
	if err != nil {
		return err
	}

	return nil
}

func (uc *TeacherUseCase) Update(id int, updatedTeacher domain.Teacher) error {
	err := uc.teacherRepo.Update(id, &updatedTeacher)
	if err != nil {
		return err
	}

	return nil
}