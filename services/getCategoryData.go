package services

import (
	"viper/models"
)

func (ps *InventoryService) GetCategory(pageNo, totalPerPage int) []*models.Category {
	category, err := models.GetDataCategory(ps.db, totalPerPage*(pageNo-1), totalPerPage)
	if err != nil {
		return nil
	}
	return category
}
