package feature

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	features := router.Group("/features")

	features.POST("", handler.Create)
	features.GET("", handler.FindAll)
	features.GET("/:id", handler.FindByID)
	features.PUT("/:id", handler.Update)
	features.DELETE("/:id", handler.Delete)
}
