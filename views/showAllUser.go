package view

import (
	"fmt"
	_ "log"
	"viper/models"
	"viper/services"
)

func (ssa *Env) GetAllUSer() {
	userService := services.NewService(ssa.db)
	user := userService.GetAllUser(1, 50)
	ViewAllUser(user)
}

func ViewAllUser(user []*models.User) {
	data := []string{"No.", "UserID", "userName"}
	fmt.Printf("\t%-10v\t%-10v\n", data[0], data[1])
	for idx, u := range user {
		fmt.Println("---------------------------------------------------------------------------------------------------------------------")
		//log.Printf("%v %v %v %v %v", p.ProductID, p.ProductDesc, p.CategoryID, p.DiscountID, p.ProductPrice)
		fmt.Printf("\t%-15v\t%-15v\t%-15v\n", idx+1, u.UserID, u.UserName)
	}
	fmt.Println("---------------------------------------------------------------------------------------------------------------------")
}
