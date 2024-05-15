package main

import (
	"log"
	"net"
	"os"

	"github.com/gin-gonic/gin"
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

	port := "55555"
	value, ok := os.LookupEnv(APIPortEnvVar)
	if ok {
		port = value
	}
	err = router.Run(net.JoinHostPort("", port))
	if err != nil {
		log.Fatal(err)
	}
}
