package routers

import (
	"net/http"

	"course-forum/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine) {
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})

	route.GET("/", func(ctx *gin.Context) { ctx.String(http.StatusOK, "hello world") })

	registerPostRoutes(route)
}

func registerPostRoutes(route *gin.Engine) {
	route.GET("/posts", controllers.GetPosts)
	route.GET("/posts/:id", controllers.FindPost)
	route.POST("/posts", controllers.CreatePost)
	route.PATCH("/posts/:id", controllers.UpdatePost)
	route.DELETE("/posts/:id", controllers.DeletePost)
}
