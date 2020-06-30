package services

import (
	"viper/models"
)

func (ps *InventoryService) GetAllUser(pageNo, totalPerPage int) []*models.User {
	user, err := models.GetAllDataUser(ps.db, totalPerPage*(pageNo-1), totalPerPage)
	if err != nil {
		return nil
	}
	return user
}
