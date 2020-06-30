package view

import (
	"fmt"
	_ "log"
	"viper/models"
	"viper/services"
)

func (ssa *Env) GetDiscount() {
	discountService := services.NewService(ssa.db)
	discount := discountService.GetDiscount(1, 50)
	ViewAllDiscount(discount)
}

func ViewAllDiscount(discount []*models.Discount) {
	data := []string{"No.", "BrandID", "BrandDesc"}
	fmt.Printf("\t%-10v\t%-10v\t%-10v\n", data[0], data[1], data[2])
	for idx, d := range discount {
		fmt.Println("---------------------------------------------------------------------------------------------------------------------")
		//log.Printf("%v %v %v %v %v", p.ProductID, p.ProductDesc, p.CategoryID, p.DiscountID, p.ProductPrice)
		fmt.Printf("\t%-15v\t%-15v\t%-15v\n", idx+1, d.DiscountID, d.DiscountDesc)
	}
	fmt.Println("---------------------------------------------------------------------------------------------------------------------")
}
