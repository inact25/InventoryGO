package services

import (
	"viper/models"
)

func (ps *InventoryService) GetDataProduct(pageNo, totalPerPage int) []*models.Product {
	products, err := models.GetDataProduct(ps.db, totalPerPage*(pageNo-1), totalPerPage)
	if err != nil {
		return nil
	}
	return products
}
