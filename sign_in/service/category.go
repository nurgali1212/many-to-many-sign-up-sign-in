package service

import (
	"sign_in/model"
	"sign_in/repository"
)

type CategoryService struct {
	repo     repository.CategoryList
	toysRepo repository.ToysList
}

func NewCategoryService(repo repository.CategoryList, toysRepo repository.ToysList) *CategoryService {
	return &CategoryService{repo: repo, toysRepo: toysRepo}
}

func (s *CategoryService) CreateCategoryService(userId, toysId int, category model.Category) (int, error) {
	_, err := s.toysRepo.GetByIdToysRepo(userId, toysId)
	if err != nil {
		// list does not exists or does not belongs to user
		return 0, err
	}

	return s.repo.CreateCategoryRepo(toysId, category)
}

func (s *CategoryService) GetAllCategoryService(userId, toysId int) ([]model.Category, error) {
	return s.repo.GetAllCategoryRepo(userId, toysId)
}

func (s *CategoryService) GetByIdCategoryService(userId, categoryId int) (model.Category, error) {
	return s.repo.GetByIdCategoryRepo(userId, categoryId)
}

func (s *CategoryService) DeleteCategoryService(userId, categoryId int) error {
	return s.repo.DeleteCategoryRepo(userId, categoryId)
}

func (s *CategoryService) UpdateCategoryService(userId, categoryId int, input model.UpdateCategoryInput) error {
	return s.repo.UpdateCategoryRepo(userId, categoryId, input)
}
