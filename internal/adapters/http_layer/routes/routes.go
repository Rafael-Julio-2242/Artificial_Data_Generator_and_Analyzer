package routes

import (
	"artificial-data-analyzer-generation/internal/adapters/http_layer/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/define-variable-types", handlers.DefineVariableTypes)
	server.POST("/data-cleaning", handlers.DataCleaning)
	server.POST("/calculate-central-trends", handlers.CalculateCentralTrends)
}
