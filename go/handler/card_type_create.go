package handler

import (
	"fmt"
	"net/http"

	"github.com/Fabfm4/bleeding-money/app"
	"github.com/gin-gonic/gin"
)

func createCardType(c *gin.Context) {
	// Simulate a list of banks
	newCardType, err := app.CreateCardType(c)
	// Return the list of banks as JSON
	if err != nil {
		fmt.Println("Error reading from input:", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": newCardType})
}
