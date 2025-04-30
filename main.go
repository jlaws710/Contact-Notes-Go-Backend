package main

import (
	"notes-system/config"
	"notes-system/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
