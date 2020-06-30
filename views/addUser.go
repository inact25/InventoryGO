package view

import (
	"log"
	"viper/services"
)

func (ssa *Env) AddNewUser(userName, userPass string) {
	userService := services.NewService(ssa.db)
	user, err := userService.AddNewUser(userName, userPass)
	if err != nil {
		log.Print(err)
	} else {
		log.Print(*user)
	}
}
