package server

import (
	"staffing-app-backend/internal/application/categoryUseCases"
	"staffing-app-backend/internal/application/skillUseCases"

	"staffing-app-backend/internal/infraestructure/repositories/memory"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine) {
	repoCategory := memory.NewCategoryRepository()
	useCaseCategory := categoryUseCases.NewCategoryUseCase(repoCategory)
	GetCategoryEndpoint := GetCategoryEndpoint(useCaseCategory)

	engine.GET("/categories/:id", GetCategoryEndpoint)

	repoSkill := memory.NewSkillRepository()
	useCaseSkill := skillUseCases.NewSkillUseCase(repoSkill)
	getSkillEndpoint := GetSkillEndpoint(useCaseSkill)
	engine.GET("/skills/:id", getSkillEndpoint)

}
