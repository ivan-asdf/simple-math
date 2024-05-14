package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ivan-asdf/simple-math/internal/api"
)

func main() {
	router := gin.Default()

	service := api.NewService()
	handler := api.NewHandler(&service)
	handler.RegisterRoutes(router)

	err := router.Run(":1234")
	if err != nil {
		log.Fatal(err)
	}
}
