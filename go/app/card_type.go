package app

import (
	"fmt"

	"github.com/Fabfm4/bleeding-money/db"
	"github.com/gin-gonic/gin"
)

func ListCardTypes() ([]db.CardType, error) {
	var card_types []db.CardType
	card_types, err := db.GetAllCardTypes()
	fmt.Println("Banks:", card_types)
	// Simulate a list of banks
	fmt.Println("err:", err)
	if err != nil {
		return nil, err
	}
	return card_types, nil
}

func GetCardType(id int) (db.CardType, error) {
	var card_type db.CardType
	card_type, err := db.GetCardTypesById(id)
	return card_type, err
}

func CreateCardType(request *gin.Context) (db.CardType, error) {
	var newCardType db.CardType
	if err := request.BindJSON(&newCardType); err != nil {
		return db.CardType{}, err
	}
	bank, err := newCardType.CreateCardType()
	if err != nil {
		return db.CardType{}, err
	}
	return bank, nil

}
