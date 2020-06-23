package services

import (
	"viper/models"
)

func (ps *InventoryService) GetDataTransaction(pageNo, totalPerPage int) []*models.Transaction {
	transaction, err := models.GetDataTransaction(ps.db, totalPerPage*(pageNo-1), totalPerPage)
	if err != nil {
		return nil
	}
	return transaction
}
