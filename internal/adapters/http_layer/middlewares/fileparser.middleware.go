package middlewares

import (
	"artificial-data-analyzer-generation/internal/domain/ports"

	"github.com/gin-gonic/gin"
)

func InjectFileParserMiddleware(fileParser ports.FileParser) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set(ports.FileParserServiceKey, fileParser)
		context.Next()
	}
}
