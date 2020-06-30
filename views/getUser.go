package view

import (
	"fmt"
	_ "log"
	"viper/models"
	"viper/services"
)

func (ssa *Env) GetUser(userName, userPass string) {
	userService := services.NewService(ssa.db)
	user := userService.GetUser(userName, userPass)
	CatchUser(user)
}

func CatchUser(user []*models.User) {
	var opt string
	for _, u := range user {
		if len(u.UserName) != 0 {
			fmt.Println("Login Successfull")
			fmt.Scan()
			MainMenu()
		}

	}
	fmt.Println("Login Failed")
	fmt.Printf("Retry ?:")
	fmt.Scan(&opt)
	Login()
}
