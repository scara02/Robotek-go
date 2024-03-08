package usecase

import (
	"go-web-robotek/services/internal/domain"
	"go-web-robotek/services/internal/repository"
)

type Group interface {
	Create(groupName string) (int, error)
	GetOne(id int) (*domain.Group, error)
	GetAll() (*[]domain.Group, error)
	Delete(id int) (int, error)
	Update(id int, updatedGroup domain.Group) error
}

type GroupUseCase struct {
	groupRepo repository.GroupRepo
}

func NewGroupUsecase(groupRepo repository.GroupRepo) *GroupUseCase {
	return &GroupUseCase{
		groupRepo: groupRepo,
	}
}

func (uc *GroupUseCase) Create(groupName string) (int, error) {
	group := domain.Group{
		GroupName: groupName,
	}

	id, err := uc.groupRepo.Create(&group)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *GroupUseCase) GetOne(id int) (*domain.Group, error) {
	group, err := uc.groupRepo.GetOne(id)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (uc *GroupUseCase) GetAll() (*[]domain.Group, error) {
	students, err := uc.groupRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return &students, nil
}

func (uc *GroupUseCase) Delete(id int) (int, error) {
	deletedID, err := uc.groupRepo.Delete(id)
	if err != nil {
		return 0, err
	}

	return deletedID, nil
}

func (uc *GroupUseCase)	Update(id int, updatedGroup domain.Group) error {
	err := uc.groupRepo.Update(id, &updatedGroup)
	if err != nil {
		return err
	}

	return nil
}