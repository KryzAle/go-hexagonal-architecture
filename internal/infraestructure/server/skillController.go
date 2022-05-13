package server

import (
	"net/http"
	"staffing-app-backend/internal/core/domain"
	"staffing-app-backend/internal/core/ports"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSkillEndpoint(useCase ports.ISkillUseCases) func(c *gin.Context) {
	return func(c *gin.Context) {
		var skill domain.Skill
		id, isExist := c.Params.Get("id")
		if isExist {
			idNumb, err := strconv.Atoi(id)
			if err != nil {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			skill, err = useCase.Get(idNumb)
			if err != nil {
				c.AbortWithStatus(http.StatusNoContent)
				return
			}
		}
		c.JSON(http.StatusOK, skill)
	}
}
