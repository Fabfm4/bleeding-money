package app

import (
	"fmt"

	"github.com/Fabfm4/bleeding-money/db"
	"github.com/gin-gonic/gin"
)

func ListCards() ([]db.Card, error) {
	var cards []db.Card
	cards, err := db.GetAllCards()
	fmt.Println("Banks:", cards)
	// Simulate a list of banks
	fmt.Println("err:", err)
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func GetCard(id int) (db.Card, error) {
	var card db.Card
	card, err := db.GetCardById(id)
	return card, err
}

func CreateCard(request *gin.Context) (db.Card, error) {
	var newCard db.Card
	if err := request.BindJSON(&newCard); err != nil {
		return db.Card{}, err
	}
	bank, err := newCard.CreateCard()
	if err != nil {
		return db.Card{}, err
	}
	return bank, nil

}
