package services

import (
	"viper/models"
)

func (ps *InventoryService) DeleteDataTransaction(transactionCode string) (*models.Transaction, error) {
	transaction := &models.Transaction{
		TransactionCode: transactionCode,
	}
	delproCode, err := models.DeleteTransaction(ps.db, transaction)
	if err != nil {
		return nil, err
	}
	transaction.TransactionCode = delproCode
	return transaction, nil
}
