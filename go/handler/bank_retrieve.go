package handler

import (
	"net/http"
	"strconv"

	"github.com/Fabfm4/bleeding-money/app"
	"github.com/gin-gonic/gin"
)

func getBank(c *gin.Context) {
	// Simulate a list of banks
	var id int
	id, _ = strconv.Atoi(c.Param("id"))
	banks, error := app.GetBank(id)
	if error != nil {
		c.Status(http.StatusNotFound)
		return
	}

	// Return the list of banks as JSON
	c.JSON(http.StatusOK, gin.H{"data": banks})
}
