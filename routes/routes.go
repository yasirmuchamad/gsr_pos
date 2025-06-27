package routes

import (
	"gsr_pos/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", controllers.ShowProduct)
	r.POST("/product", controllers.AddProduct)
}
