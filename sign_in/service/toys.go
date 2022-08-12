package service

import (
	"sign_in/model"
	"sign_in/repository"
)

type ToysService struct {
	repo repository.ToysList
}

func NewToysService(repo repository.ToysList) *ToysService {
	return &ToysService{repo: repo}
}

func (s *ToysService) CreateToysService(userId int, toys model.Toys) (int, error) {
	return s.repo.CreateToysRepo(userId, toys)
}

func (s *ToysService) GetAllToysService(userId int) ([]model.Toys, error) {
	return s.repo.GetAllToysRepo(userId)
}

func (s *ToysService) GetByIdToysService(userId, toysId int) (model.Toys, error) {
	return s.repo.GetByIdToysRepo(userId, toysId)
}

func (s *ToysService) DeleteToysService(userId, toysId int) error {
	return s.repo.DeleteToysRepo(userId, toysId)
}

func (s *ToysService) UpdateToysService(userId, toysId int, input model.UpdateToysInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.UpdateToysRepo(userId, toysId, input)
}
