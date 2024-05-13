package main

import (
	"log"

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
	err := router.Run(":1234")
	if err != nil {
		log.Fatal(err)
	}
}
