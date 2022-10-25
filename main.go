package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"environment": fmt.Sprintf("%+v", os.Environ()),
			"headers":     fmt.Sprintf("%+v", c.Request.Header),
		})
	})

	r.GET("/live", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK++")
	})

	r.GET("/ready", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK++")
	})

	r.Run()
}
