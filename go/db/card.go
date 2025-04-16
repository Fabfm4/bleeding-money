package db

import (
	"errors"
	"time"
)

type Card struct {
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name"`
	Last4digits  string `json:"last_4_digits"`
	Card_type_id int    `json:"card_type_id"`
	Bank_id      int    `json:"bank_id"`
}

var dummyListCards = []Card{
	{
		CreatedAt:    "2023-10-01",
		UpdatedAt:    "2023-10-01",
		ID:           1,
		Name:         "Card A",
		Last4digits:  "1234",
		Card_type_id: 1,
		Bank_id:      1,
	},
	{
		CreatedAt:    "2023-10-01",
		UpdatedAt:    "2023-10-01",
		ID:           2,
		Name:         "Card B",
		Last4digits:  "5678",
		Card_type_id: 2,
		Bank_id:      2,
	},
}

func GetAllCards() ([]Card, error) {
	return dummyListCards, nil
}

func GetCardById(id int) (Card, error) {
	var card_found Card

	for _, card := range dummyListCards {
		if card.ID == id {
			card_found = card
			break
		}
	}

	if card_found.ID == 0 {
		return card_found, errors.New("bank not found")
	}

	return card_found, nil
}

func (card Card) CreateCard() (Card, error) {
	card.ID = len(dummyListBanks) + 1
	card.CreatedAt = time.Now().UTC().String()
	dummyListCards = append(dummyListCards, card)
	return card, nil
}
