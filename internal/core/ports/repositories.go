package ports

import "staffing-app-backend/internal/core/domain"

type ICategoryRepository interface {
	Get(id int) (domain.Category, error)
	Save(category domain.Category) error
}
type ISkillRepository interface {
	Get(id int) (domain.Skill, error)
	Save(category domain.Skill) error
}
