package view

import (
	"log"
	"viper/services"
)

func (ssa *Env) AddNewProduct(productID, productDesc, categoryID, productPrice, discountID, brandCode string) {
	productService := services.NewService(ssa.db)
	prod, err := productService.AddNewProduct(productID, productDesc, categoryID, productPrice, discountID, brandCode)
	if err != nil {
		log.Print(err)
	} else {
		log.Print(*prod)
	}
}
