package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ivan-asdf/simple-math/api"
)

func main() {
	// input := "What iS 4 divided by 1?"

	router := gin.Default()
	router.POST("evaluate", api.Evaluate)
	router.POST("validate", api.Validate)
	// router.GET("/", func(c *gin.Context) {
	//   c.String(http.StatusOK, "GET method\n")
	// })
	router.Run(":1234")
}
