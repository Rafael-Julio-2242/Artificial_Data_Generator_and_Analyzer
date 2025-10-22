package middlewares

import (
	"artificial-data-analyzer-generation/internal/domain/services"

	"github.com/gin-gonic/gin"
)

func InjectCalculateFrequenciesMiddleware(calculateFrequenciesService services.CalculateFrequencies) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set(services.CalculateFrequenciesServiceKey, calculateFrequenciesService)
		context.Next()
	}
}
