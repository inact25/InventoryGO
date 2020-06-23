package view

import (
	"log"
	"viper/services"
)

func (ssa *Env) DeleteDataTransaction(transactionCode string) {
	transactionService := services.NewService(ssa.db)
	prod, err := transactionService.DeleteDataTransaction(transactionCode)
	if err != nil {
		log.Print(err)
	} else {
		log.Print(*prod)
	}
}
