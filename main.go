package main

import (
	"gsr_pos/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	config.ConnectDatabase()

	routes.SetupRoutes(r)

	r.Run(":8000")
}
