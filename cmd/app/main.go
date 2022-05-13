package main

import (
	"log"
	"staffing-app-backend/internal/infraestructure/server"
	"staffing-app-backend/shared/utils/envvar"

	"github.com/gin-gonic/gin"
)

func main() {
	PORT := envvar.ApiPort()

	engine := gin.Default()
	server.RegisterRouter(engine)
	log.Fatal(engine.Run(":" + PORT))
}
