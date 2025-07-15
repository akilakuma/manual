package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	var r *gin.Engine
	r = gin.Default()
	meow := r.Group("/")

	meow.GET("hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello man !")
		return
	})
	r.Run(":2266")
}
