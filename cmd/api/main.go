package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ivan-asdf/simple-math/cmd/common"
	"github.com/ivan-asdf/simple-math/internal/api"
	"github.com/joho/godotenv"
)

const APIPortEnvVar = "API_PORT"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()

	service := api.NewService()
	handler := api.NewHandler(&service)
	handler.RegisterRoutes(router)

	port := common.DefaultPort
	value, ok := os.LookupEnv(APIPortEnvVar)
	if ok {
		port = ":" + value
	}
	err = router.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
