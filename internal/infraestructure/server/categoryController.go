package server

import (
	"net/http"
	"staffing-app-backend/internal/core/domain"
	"staffing-app-backend/internal/core/ports"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCategoryEndpoint(useCase ports.ICategoryUseCases) func(c *gin.Context) {
	return func(c *gin.Context) {
		var category domain.Category
		id, isExist := c.Params.Get("id")
		if isExist {
			idNumb, err := strconv.Atoi(id)
			if err != nil {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			category, err = useCase.Get(idNumb)
			if err != nil {
				c.AbortWithStatus(http.StatusNoContent)
				return
			}
		}
		c.JSON(http.StatusOK, category)
	}
}
