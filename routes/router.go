package routes

import (
	controller "notes-system/controller"
	"notes-system/middleware"
	"notes-system/util"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/login", func(c *gin.Context) {
		token := util.GenerateToken()
		util.StoreToken(token)
		c.JSON(200, gin.H{"token": token})
	})

	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/contacts", controller.GetContacts)
		auth.POST("/contacts", controller.CreateContact)
		auth.PUT("/contacts/:id", controller.UpdateContact)
		auth.DELETE("/contacts/:id", controller.DeleteContact)
		auth.GET("/contacts/:id/notes", controller.GetNotes)
		auth.POST("/contacts/:id/notes", controller.CreateNote)
	}
}
