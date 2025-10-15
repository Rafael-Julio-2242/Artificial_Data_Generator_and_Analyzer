package middlewares

import (
	"artificial-data-analyzer-generation/internal/domain/services"

	"github.com/gin-gonic/gin"
)

func InjectDataCleaningMiddleware(dataCleaningService services.DataCleaningService) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set(services.DataCleaningServiceKey, dataCleaningService)
		context.Next()
	}
}
