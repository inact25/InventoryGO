package services

import (
	"database/sql"
	"viper/models"
)

type TransactionService struct {
	db *sql.DB
}

func ViewAllTransactionService(db *sql.DB) *TransactionService {
	return &TransactionService{db}
}

func (ps *TransactionService) GetDataTransaction(pageNo, totalPerPage int) []*models.Transaction {
	transaction, err := models.GetDataTransaction(ps.db, totalPerPage*(pageNo-1), totalPerPage)
	if err != nil {
		return nil
	}
	return transaction
}
