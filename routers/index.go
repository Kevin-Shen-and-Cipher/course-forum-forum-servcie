package routers

import (
	"net/http"

	"course-forum/controllers"

	docs "course-forum/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine) {
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})

	registerSwaggerRoutes(route)

	baseRoute := route.Group("/api/v1")
	registerPostRoutes(baseRoute)
	registerTagRoutes(baseRoute)
}

func registerSwaggerRoutes(route *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func registerPostRoutes(route *gin.RouterGroup) {
	postRoute := route.Group("posts")

	postRoute.GET("", controllers.GetPosts)
	postRoute.GET(":id", controllers.FindPost)
	postRoute.POST("", controllers.CreatePost)
	postRoute.PATCH(":id", controllers.UpdatePost)
	postRoute.DELETE(":id", controllers.DeletePost)
}

func registerTagRoutes(route *gin.RouterGroup) {
	tagRoute := route.Group("tags")

	tagRoute.GET("", controllers.GetTags)
	tagRoute.GET(":id", controllers.FindTag)
	tagRoute.POST("", controllers.CreateTag)
	tagRoute.PATCH(":id", controllers.UpdateTag)
	tagRoute.DELETE(":id", controllers.DeleteTag)
}
