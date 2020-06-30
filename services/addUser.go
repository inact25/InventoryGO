package services

import (
	"viper/models"
)

func (ps *InventoryService) AddNewUser(userName, userPass string) (*models.User, error) {
	user := &models.User{
		UserName: userName,
		UserPass: userPass,
	}
	id, err := models.AddNewUser(ps.db, user)
	if err != nil {
		return nil, err
	}
	user.UserID = id
	return user, nil
}
