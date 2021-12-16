package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	defaultUrl := "http://image-cache.ecf-cluster"
	router := gin.Default()
	router.GET("/*imagePath", func(c *gin.Context) {
		imagePath := c.Param("imagePath")
		fmt.Println(imagePath)
		url := defaultUrl + imagePath
		c.JSON(http.StatusOK, gin.H{
			"image_url": url,
		})
	})

	_ = router.Run(":9999")
}
