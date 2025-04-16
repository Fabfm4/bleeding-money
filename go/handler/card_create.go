package handler

import (
	"fmt"
	"net/http"

	"github.com/Fabfm4/bleeding-money/app"
	"github.com/gin-gonic/gin"
)

func createCard(c *gin.Context) {
	// Simulate a list of banks
	newCard, err := app.CreateCard(c)
	// Return the list of banks as JSON
	if err != nil {
		fmt.Println("Error reading from input:", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": newCard})

}
