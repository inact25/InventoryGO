package view

import (
	"log"
	"viper/services"
)

func (ssa *Env) DeleteDataProduct(productID string) {
	productService := services.NewService(ssa.db)
	prod, err := productService.DeleteDataProduct(productID)
	if err != nil {
		log.Print(err)
	} else {
		log.Print(*prod)
	}
}
