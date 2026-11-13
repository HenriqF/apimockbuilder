package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func config_endpoints(eng *gin.Engine) {
	eng.GET("/lista", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	eng.POST("/novo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	eng.POST("/config", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	eng.PUT("/config", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})


}

func Start() {
	eng := gin.Default()

	config_endpoints(eng)

	eng.Run(":3000")
}
