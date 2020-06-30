package view

import (
	"fmt"
	_ "log"
	"viper/models"
	"viper/services"
)

func (ssa *Env) GetBrand() {
	brandService := services.NewService(ssa.db)
	brand := brandService.GetBrand(1, 50)
	ViewAllBrand(brand)
}

func ViewAllBrand(brand []*models.Brand) {
	data := []string{"No.", "BrandID", "BrandDesc"}
	fmt.Printf("\t%-10v\t%-10v\t%-10v\n", data[0], data[1], data[2])
	for idx, b := range brand {
		fmt.Println("---------------------------------------------------------------------------------------------------------------------")
		//log.Printf("%v %v %v %v %v", p.ProductID, p.ProductDesc, p.CategoryID, p.DiscountID, p.ProductPrice)
		fmt.Printf("\t%-15v\t%-15v\t%-15v\n", idx+1, b.BrandID, b.BrandDesc)
	}
	fmt.Println("---------------------------------------------------------------------------------------------------------------------")
}
