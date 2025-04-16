package handler

import (
	"net/http"

	"github.com/Fabfm4/bleeding-money/app"
	"github.com/gin-gonic/gin"
)

func listCardTypes(c *gin.Context) {
	// Simulate a list of banks
	cardTypes, _ := app.ListCardTypes()
	// Return the list of banks as JSON

	c.JSON(http.StatusOK, gin.H{"data": cardTypes})

}
