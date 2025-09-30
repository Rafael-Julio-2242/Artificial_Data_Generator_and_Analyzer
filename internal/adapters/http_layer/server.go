package httplayer

import (
	"artificial-data-analyzer-generation/internal/adapters/http_layer/routes"

	"github.com/gin-gonic/gin"
)

func GetHttpServer(middlewares ...gin.HandlerFunc) *gin.Engine {
	server := gin.Default()
	for _, middleware := range middlewares {
		server.Use(middleware)
	}
	routes.RegisterRoutes(server)
	return server
}
