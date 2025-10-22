package middlewares

import (
	"artificial-data-analyzer-generation/internal/domain/ports"

	"github.com/gin-gonic/gin"
)

func InjectCSVGeneratorMiddleware(csvGenerator ports.CSVGenerator) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set(ports.CSVGeneratorServiceKey, csvGenerator)
		context.Next()
	}
}
