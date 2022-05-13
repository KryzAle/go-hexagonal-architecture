package memory

import (
	"staffing-app-backend/internal/core/domain"
	"sync"
)

var (
	onceSkillsRepo     sync.Once
	instanceSkillsRepo *skillsRepository
)

type skillsRepository struct {
	dataMemory map[int]domain.Skill
	lastId     int
}

func NewSkillRepository() *skillsRepository {
	onceSkillsRepo.Do(func() {
		instanceSkillsRepo = &skillsRepository{
			dataMemory: map[int]domain.Skill{},
			lastId:     0,
		}
	})
	return instanceSkillsRepo
}

func (r *skillsRepository) Get(id int) (domain.Skill, error) {
	skill := r.dataMemory[id]
	return skill, nil
}

func (r *skillsRepository) Save(skill domain.Skill) error {
	r.dataMemory[r.lastId] = skill
	r.lastId++
	return nil
}
