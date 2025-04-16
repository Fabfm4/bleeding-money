package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var blueprint = "/api/v1"

func LoadHandlers(router *gin.Engine) {
	router.GET(fmt.Sprintf("%s%s", blueprint, "/banks"), listBanks)
	router.POST(fmt.Sprintf("%s%s", blueprint, "/banks"), createBank)
	router.GET(fmt.Sprintf("%s%s", blueprint, "/banks/:id"), getBank)

	router.GET(fmt.Sprintf("%s%s", blueprint, "/cards"), listCards)
	router.POST(fmt.Sprintf("%s%s", blueprint, "/cards"), createCard)
	router.GET(fmt.Sprintf("%s%s", blueprint, "/cards/:id"), getCard)

	router.GET(fmt.Sprintf("%s%s", blueprint, "/card-types"), listCardTypes)
	router.POST(fmt.Sprintf("%s%s", blueprint, "/card-types"), createCardType)
	router.GET(fmt.Sprintf("%s%s", blueprint, "/card-types/:id"), getCardType)
}
