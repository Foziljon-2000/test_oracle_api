package main

import (
	"github.com/Foziljon-2000/test_oracle_api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/", handlers.JsonGan)

	router.Run(":4001")
}
