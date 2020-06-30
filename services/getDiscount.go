package services

import (
	"viper/models"
)

func (ps *InventoryService) GetDiscount(pageNo, totalPerPage int) []*models.Discount {
	discount, err := models.GetDataDiscount(ps.db, totalPerPage*(pageNo-1), totalPerPage)
	if err != nil {
		return nil
	}
	return discount
}
