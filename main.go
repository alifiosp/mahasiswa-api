package main

import (
	"net/http"

	"github.com/alifiosp/mahasiswa-api/models"
	"github.com/alifiosp/mahasiswa-api/models/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Data": "Home Page",
		})
	})

	r.GET("/mahasiswa", controllers.Tampil)
	r.POST("/mahasiswa", controllers.MahasiswaTambah)

	r.Run()
}
