package services

import (
	"viper/models"
)

func (ps *InventoryService) AddNewTransaction(transactionCode, transactionDate, productDesc, transactionQTY string) (*models.Transaction, error) {
	transaction := &models.Transaction{
		TransactionCode: transactionCode,
		TransactionDate: transactionDate,
		ProductDesc:     productDesc,
		TransactionQTY:  transactionQTY,
	}
	id, err := models.AddNewTransaction(ps.db, transaction)
	if err != nil {
		return nil, err
	}
	transaction.TransactionCode = id
	return transaction, nil
}
