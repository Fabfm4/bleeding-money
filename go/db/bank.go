package db

import (
	"errors"
	"time"
)

type Bank struct {
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
}

var dummyListBanks = []Bank{
	{
		CreatedAt: "2023-10-01",
		UpdatedAt: "2023-10-01",
		ID:        1,
		Name:      "Bank A",
	},
	{
		CreatedAt: "2023-10-01",
		UpdatedAt: "2023-10-01",
		ID:        2,
		Name:      "Bank B",
	},
}

func GetAllBanks() ([]Bank, error) {
	return dummyListBanks, nil
}

func GetBanksById(id int) (Bank, error) {
	var bank_found Bank

	for _, bank := range dummyListBanks {
		if bank.ID == id {
			bank_found = bank
			break
		}
	}

	if bank_found.ID == 0 {
		return bank_found, errors.New("bank not found")
	}

	return bank_found, nil
}

func (bank Bank) CreateBank() (Bank, error) {
	bank.ID = len(dummyListBanks) + 1
	bank.CreatedAt = time.Now().UTC().String()
	dummyListBanks = append(dummyListBanks, bank)
	return bank, nil
}
