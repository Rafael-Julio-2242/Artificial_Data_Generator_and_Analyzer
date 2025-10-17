package middlewares

import (
	"artificial-data-analyzer-generation/internal/domain/services"

	"github.com/gin-gonic/gin"
)

func InjectCalculateCentralTrendsMiddleware(calculateCentralTrendsService services.CalculateCentralTrends) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set(services.CalculateCentralTrendsServiceKey, calculateCentralTrendsService)
		context.Next()
	}
}
