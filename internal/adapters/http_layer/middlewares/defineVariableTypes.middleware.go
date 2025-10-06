package middlewares

import (
	"artificial-data-analyzer-generation/internal/domain/services"

	"github.com/gin-gonic/gin"
)

func InjectDefineVariableTypesMiddleware(defineVariableTypesService services.DefineVariableTypesService) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set(services.DefineVariableTypesServiceKey, defineVariableTypesService)
		context.Next()
	}
}
