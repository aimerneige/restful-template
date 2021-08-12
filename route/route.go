package route

import (
	"restful-template/controllers"
	"restful-template/middleware"

	"github.com/gin-gonic/gin"
)

func RouteCollection(r *gin.Engine) *gin.Engine {

	r.GET("/", controllers.Index)

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	articleRoutes := r.Group("/article")
	articleRoutes.Use(middleware.AuthMiddleware())
	articleRoutes.GET("", controllers.GetArticle)
	articleRoutes.DELETE("", controllers.DeleteArticle)
	articleRoutes.Use(middleware.ArticleMiddleware())
	articleRoutes.POST("", controllers.CreateArticle)
	articleRoutes.PUT("", controllers.UpdateArticle)

	return r
}
