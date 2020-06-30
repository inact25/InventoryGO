package services

import (
	"viper/models"
)

func (ps *InventoryService) GetUser(userName, userPass string) []*models.User {
	user, err := models.GetDataUser(ps.db, userName, userPass)
	if err != nil {
		return nil
	}
	return user
}
