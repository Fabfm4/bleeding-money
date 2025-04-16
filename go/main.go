package main

import (
	"github.com/Fabfm4/bleeding-money/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	handler.LoadHandlers(router)
	router.Run("localhost:8080")
}
