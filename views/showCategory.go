package view

import (
	"fmt"
	_ "log"
	"viper/models"
	"viper/services"
)

func (ssa *Env) GetCategory() {
	categoryService := services.NewService(ssa.db)
	category := categoryService.GetCategory(1, 50)
	ViewAllCategory(category)
}

func ViewAllCategory(category []*models.Category) {
	data := []string{"No.", "CategoryID", "categoryDesc"}
	fmt.Printf("\t%-10v\t%-10v\t%-10v\n", data[0], data[1], data[2])
	for idx, c := range category {
		fmt.Println("---------------------------------------------------------------------------------------------------------------------")
		//log.Printf("%v %v %v %v %v", p.ProductID, p.ProductDesc, p.CategoryID, p.DiscountID, p.ProductPrice)
		fmt.Printf("\t%-15v\t%-15v\t%-15v\n", idx+1, c.CategoryID, c.CategoryDesc)
	}
	fmt.Println("---------------------------------------------------------------------------------------------------------------------")
}
