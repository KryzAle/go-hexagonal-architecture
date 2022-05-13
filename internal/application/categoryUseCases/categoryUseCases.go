package categoryUseCases

import (
	"staffing-app-backend/internal/core/domain"
	"staffing-app-backend/internal/core/ports"
)

type categoryUseCase struct {
	repo ports.ICategoryRepository
}

func NewCategoryUseCase(repo ports.ICategoryRepository) *categoryUseCase {
	return &categoryUseCase{
		repo: repo,
	}
}

func (s *categoryUseCase) Get(id int) (domain.Category, error) {
	return s.repo.Get(id)
}

func (s *categoryUseCase) Create(category domain.Category) error {
	return s.repo.Save(category)
}
