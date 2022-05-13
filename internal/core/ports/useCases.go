package ports

import "staffing-app-backend/internal/core/domain"

type ICategoryUseCases interface {
	Get(id int) (domain.Category, error)
	Create(category domain.Category) error
}
type ISkillUseCases interface {
	Get(id int) (domain.Skill, error)
	Create(category domain.Skill) error
}