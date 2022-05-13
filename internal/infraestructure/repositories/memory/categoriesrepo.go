package memory

import (
	"staffing-app-backend/internal/core/domain"
	"sync"
)

var (
	onceCategorysRepo     sync.Once
	instanceCategorysRepo *repository
)

type repository struct {
	dataMemory map[int]domain.Category
	lastId     int
}

func NewCategoryRepository() *repository {
	onceCategorysRepo.Do(func() {
		instanceCategorysRepo = &repository{
			dataMemory: map[int]domain.Category{},
			lastId:     0,
		}
	})
	return instanceCategorysRepo
}

func (r *repository) Get(id int) (domain.Category, error) {
	category := r.dataMemory[id]
	return category, nil
}

func (r *repository) Save(category domain.Category) error {
	r.dataMemory[r.lastId] = category
	r.lastId++
	return nil
}
