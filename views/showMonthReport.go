package view

import (
	"fmt"
	_ "log"
	"time"
	"viper/models"
	"viper/services"
)

func (ssa *Env) GetMonthReport() {
	currentTime := time.Now()
	time := currentTime.Format("01-02")
	getTime := fmt.Sprintf("%%" + time + "%%")

	reportService := services.NewService(ssa.db)
	report := reportService.GetMonthReport(getTime, 1, 50)
	ViewMonthReport(report)
}

func ViewMonthReport(report []*models.Report) {
	data := []string{"No.", "Transaction Date", "Code", "Name", "QTY", "Price", "subTotal"}
	fmt.Printf("\t%-10v\t%-10v\t%-10v\t%-10v\t%-10v\t%-10v\t%-10v\n", data[0], data[1], data[2], data[3], data[4], data[5], data[6])
	for idx, r := range report {
		fmt.Println("---------------------------------------------------------------------------------------------------------------------")
		//log.Printf("%v %v %v %v %v", p.ProductID, p.ProductDesc, p.CategoryID, p.DiscountID, p.ProductPrice)
		fmt.Printf("\t%-15v\t%-15v\t%-15v\t%-15v\t%-15v\t%-15v\t%-15v\n", idx+1, r.TransactionDate, r.TransactionCode, r.ProductDesc, r.TransactionQTY, r.ProductPrice, r.SubTotal)
	}
	fmt.Println("---------------------------------------------------------------------------------------------------------------------")
}
