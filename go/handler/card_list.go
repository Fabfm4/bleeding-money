package handler

import (
	"net/http"

	"github.com/Fabfm4/bleeding-money/app"
	"github.com/gin-gonic/gin"
)

func listCards(c *gin.Context) {
	// Simulate a list of banks
	cards, _ := app.ListCards()
	// Return the list of banks as JSON

	c.JSON(http.StatusOK, gin.H{"data": cards})

}
