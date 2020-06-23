package view

import (
	"log"
	"time"
	"viper/services"
)

func (ssa *Env) AddNewTransaction(transactionCode, productDesc, transactionQTY string) {
	currentTime := time.Now()
	transactionDate := currentTime.Format("2006.01.02 15:04:05")
	transactionService := services.NewService(ssa.db)
	trans, err := transactionService.AddNewTransaction(transactionCode, transactionDate, productDesc, transactionQTY)

	if err != nil {
		log.Print(err)
	} else {
		log.Print(*trans)
	}
}
