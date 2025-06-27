package controllers

import (
	"gsr_pos/config"
	"gsr_pos/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowProduct(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.HTML(http.StatusOK, "index.html", gin.H{"products": products})
}

func AddProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBind(&product); err == nil {
		config.DB.Create(&product)
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
