package view

import (
	"fmt"
	_ "log"
	"viper/models"
)

func ViewAllTransaction(transaction []*models.Transaction) {
	data := []string{"No.", "Code", "Date", "Brand", "Product", "QTY"}
	fmt.Printf("\t%-10v\t%-10v\t%-25v\t%-10v\t%-10v\t%-10v\n", data[0], data[1], data[2], data[3], data[4], data[5])
	for idx, p := range transaction {
		fmt.Println("---------------------------------------------------------------------------------------------------------------------")
		//log.Printf("%v %v %v %v %v", p.ProductID, p.ProductDesc, p.CategoryID, p.DiscountID, p.ProductPrice)
		fmt.Printf("\t%-10v\t%-10v\t%-25v\t%-10v\t%-10v\t%-10v\n", idx+1, p.TransactionCode, p.TransactionDate, p.BrandDesc, p.ProductDesc, p.TransactionQTY)
	}
	fmt.Println("---------------------------------------------------------------------------------------------------------------------")
}
