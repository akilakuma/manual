package main

import "github.com/gin-gonic/gin"

func main() {

	// default start
	// gin.Default()

	// release strart

	gin.SetMode(gin.ReleaseMode)
	gin.New()

}


	// 健康偵測B
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	// 以下才開始加log
	r.Use(gin.Logger())