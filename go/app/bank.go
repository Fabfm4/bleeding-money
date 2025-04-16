package app

import (
	"fmt"

	"github.com/Fabfm4/bleeding-money/db"
	"github.com/gin-gonic/gin"
)

func ListBanks() ([]db.Bank, error) {
	var banks []db.Bank
	banks, err := db.GetAllBanks()
	fmt.Println("Banks:", banks)
	// Simulate a list of banks
	fmt.Println("err:", err)
	if err != nil {
		return nil, err
	}
	return banks, nil
}

func GetBank(id int) (db.Bank, error) {
	var bank db.Bank
	bank, err := db.GetBanksById(id)
	return bank, err
}

func CreateBank(request *gin.Context) (db.Bank, error) {
	var newBank db.Bank
	if err := request.BindJSON(&newBank); err != nil {
		return db.Bank{}, err
	}
	bank, err := newBank.CreateBank()
	if err != nil {
		return db.Bank{}, err
	}
	return bank, nil

}
