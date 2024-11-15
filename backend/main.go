package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET  EMPLOYEE METHOD",
		})
	})

	//POST METHOD
	r.POST("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST EMPLOYEE METHOD",
		})
	})

	//PUT METHOD
	r.PUT("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT EMPLOYEE METHOD",
		})
	})

	//DELETE METHOD
	r.DELETE("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE EMPLOYEE METHOD",
		})
	})
		r.Run(":8081") // listen and serve on 0.0.0.0:8081 (for windows "localhost:8080")
}
