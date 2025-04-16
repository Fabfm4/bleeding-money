package db

import (
	"errors"
	"time"
)

type CardType struct {
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
}

var dummyListCardTypes = []CardType{
	{
		CreatedAt: "2023-10-01",
		UpdatedAt: "2023-10-01",
		ID:        1,
		Name:      "DEBIT",
	},
	{
		CreatedAt: "2023-10-01",
		UpdatedAt: "2023-10-01",
		ID:        2,
		Name:      "CREDIT",
	},
}

func GetAllCardTypes() ([]CardType, error) {
	return dummyListCardTypes, nil
}

func GetCardTypesById(id int) (CardType, error) {
	var creadit_type_found CardType

	for _, creadit_type := range dummyListCardTypes {
		if creadit_type.ID == id {
			creadit_type_found = creadit_type
			break
		}
	}

	if creadit_type_found.ID == 0 {
		return creadit_type_found, errors.New("bank not found")
	}

	return creadit_type_found, nil
}

func (card_type CardType) CreateCardType() (CardType, error) {
	card_type.ID = len(dummyListBanks) + 1
	card_type.CreatedAt = time.Now().UTC().String()
	dummyListCardTypes = append(dummyListCardTypes, card_type)
	return card_type, nil
}
