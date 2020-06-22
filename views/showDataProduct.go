package view

import (
	"fmt"
	_ "log"
	"viper/models"
)

func ViewAllProduct(product []*models.Product) {
	data := []string{"No.", "Product ID", "Brand", "Name", "Type", "Price", "Discount"}
	fmt.Printf("\t%-10v\t%-10v\t%-10v\t%-10v\t%-10v\t%-10v\t%-10v\n", data[0], data[1], data[2], data[3], data[4], data[5], data[6])
	for idx, p := range product {
		fmt.Println("---------------------------------------------------------------------------------------------------------------------")
		//log.Printf("%v %v %v %v %v", p.ProductID, p.ProductDesc, p.CategoryID, p.DiscountID, p.ProductPrice)
		fmt.Printf("\t%-15v\t%-15v\t%-15v\t%-15v\t%-15v\t%-15v\t%-15v\n", idx+1, p.ProductID, p.BrandDesc, p.ProductDesc, p.CategoryID, p.DiscountID, p.ProductPrice)
	}
	fmt.Println("---------------------------------------------------------------------------------------------------------------------")
}
