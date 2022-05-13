package memory

import (
	"staffing-app-backend/internal/core/domain"
	"sync"
)

var (
	onceCategorysRepo      sync.Once
	instanceCategoriesRepo *categoriesRepository
)

type categoriesRepository struct {
	dataMemory map[int]domain.Category
	lastId     int
}

func NewCategoryRepository() *categoriesRepository {
	onceCategorysRepo.Do(func() {
		instanceCategoriesRepo = &categoriesRepository{
			dataMemory: map[int]domain.Category{},
			lastId:     0,
		}
	})
	return instanceCategoriesRepo
}

func (r *categoriesRepository) Get(id int) (domain.Category, error) {
	category := r.dataMemory[id]
	return category, nil
}

func (r *categoriesRepository) Save(category domain.Category) error {
	r.dataMemory[r.lastId] = category
	r.lastId++
	return nil
}
