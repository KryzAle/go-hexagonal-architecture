package skillUseCases

import (
	"staffing-app-backend/internal/core/domain"
	"staffing-app-backend/internal/core/ports"
)

type skillUseCase struct {
	repo ports.ISkillRepository
}

func NewSkillUseCase(repo ports.ISkillRepository) *skillUseCase {
	return &skillUseCase{
		repo: repo,
	}
}

func (s *skillUseCase) Get(id int) (domain.Skill, error) {
	return s.repo.Get(id)
}

func (s *skillUseCase) Create(skill domain.Skill) error {
	return s.repo.Save(skill)
}
