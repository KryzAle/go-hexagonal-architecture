package server

import (
	"staffing-app-backend/internal/application/categoryUseCases"
	"staffing-app-backend/internal/infraestructure/repositories/memory"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine) {
	repoCategory := memory.NewCategoryRepository()
	servCategory := categoryUseCases.NewCategoryUseCase(repoCategory)
	getEndpoint := GetCategoryEndpoint(servCategory)

	engine.GET("/categories/:id", getEndpoint)
}
