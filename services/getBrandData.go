package services

import (
	"viper/models"
)

func (ps *InventoryService) GetBrand(pageNo, totalPerPage int) []*models.Brand {
	brand, err := models.GetDataBrand(ps.db, totalPerPage*(pageNo-1), totalPerPage)
	if err != nil {
		return nil
	}
	return brand
}
