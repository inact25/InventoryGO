package view

import (
	"fmt"
	"viper/config"
	"viper/tools"
)

var conf = config.NewAppConfig()

func MainMenu() {
	var choose string

	//id := 1
	menu := []string{"Product's Data", "Product's Transaction", "Product's Report"}
	fmt.Println("-----------------------------------------------")
	fmt.Println("Go Inventory")
	fmt.Println("-----------------------------------------------")
	fmt.Printf("\t%-10s\t%-10s\n", "No.", "Detail")
	for i, val := range menu {
		fmt.Println("-----------------------------------------------")
		fmt.Printf("\t%-10v\t%-10v\n", i+1, val)
	}
	fmt.Println("-----------------------------------------------")
	fmt.Printf("What you Wanna do ? : ")
	fmt.Scan(&choose)
	fmt.Println()
	chooseOption("1", choose)
}

func chooseOption(id, option string) {
	var choose string
	switch id {
	case "1":
		switch option {
		case "1":
			menu := []string{"Add Product", "Delete Product", "Show All Product"}
			fmt.Println("-----------------------------------------------")
			fmt.Println("Product's Data")
			fmt.Println("-----------------------------------------------")
			fmt.Printf("\t%-10s\t%-10s\n", "No.", "Detail")
			for i, val := range menu {
				fmt.Println("-----------------------------------------------")
				fmt.Printf("\t%-10v\t%-10v\n", i+1, val)
			}
			fmt.Println("-----------------------------------------------")
			fmt.Printf("What you Wanna do ? : ")
			fmt.Scan(&choose)
			fmt.Println()
			chooseOption("1.1", choose)
		case "2":
			menu := []string{"Add Transaction", "Delete Transaction", "Show All Transaction"}
			fmt.Println("-----------------------------------------------")
			fmt.Println("Product's Transaction")
			fmt.Println("-----------------------------------------------")
			fmt.Printf("\t%-10s\t%-10s\n", "No.", "Detail")
			for i, val := range menu {
				fmt.Println("-----------------------------------------------")
				fmt.Printf("\t%-10v\t%-10v\n", i+1, val)
			}
			fmt.Println("-----------------------------------------------")
			fmt.Printf("What you Wanna do ? : ")
			fmt.Scan(&choose)
			fmt.Println()
			chooseOption("1.2", choose)
		case "3":
			menu := []string{"Daily Report", "Monthly Report", "Show All Report"}
			fmt.Println("-----------------------------------------------")
			fmt.Println("Product's Report")
			fmt.Println("-----------------------------------------------")
			fmt.Printf("\t%-10s\t%-10s\n", "No.", "Detail")
			for i, val := range menu {
				fmt.Println("-----------------------------------------------")
				fmt.Printf("\t%-10v\t%-10v\n", i+1, val)
			}
			fmt.Println("-----------------------------------------------")
			fmt.Printf("What you Wanna do ? : ")
			fmt.Scan(&choose)
			fmt.Println()
			chooseOption("1.3", choose)
		}

	case "1.1":
		switch option {
		case "1":
			var init, productCode, productDesc, categoryID, productPrice, discountID, brandCode string
			fmt.Println("-----------------------------------------------")
			fmt.Println("Add Product's Data")
			fmt.Println("-----------------------------------------------")
			tools.Input("", &init)
			tools.Input("Insert Product ID : ", &productCode)
			tools.Input("Insert Product Name : ", &productDesc)
			tools.Input("Insert Brand Code : ", &brandCode)
			tools.Input("Insert Category ID : ", &categoryID)
			tools.Input("Insert Product Price : ", &productPrice)
			tools.Input("Insert Discount ID : ", &discountID)
			Execute(conf).AddNewProduct(productCode, productDesc, categoryID, productPrice, discountID, brandCode)
		case "2":
			var init, productCode string
			fmt.Println("-----------------------------------------------")
			fmt.Println("Delete Product's Data")
			fmt.Println("-----------------------------------------------")
			tools.Input("", &init)
			tools.Input("Insert Product ID to Delete: ", &productCode)
			Execute(conf).DeleteDataProduct(productCode)
		case "3":
			fmt.Println("-----------------------------------------------")
			fmt.Println("Show All Product's Data")
			fmt.Println("-----------------------------------------------")
			Execute(conf).GetDataProduct()
		}
	case "1.2":
		switch option {
		case "1":
			var init, transactionCode, productId, transactionQTY string
			fmt.Println("-----------------------------------------------")
			fmt.Println("Add Transaction's Data")
			fmt.Println("-----------------------------------------------")
			tools.Input("", &init)
			tools.Input("Insert Transaction ID : ", &transactionCode)
			tools.Input("Insert Product ID : ", &productId)
			tools.Input("Insert Quantity : ", &transactionQTY)
			Execute(conf).AddNewTransaction(transactionCode, productId, transactionQTY)
		case "2":
			var init, transactionCode string
			fmt.Println("-----------------------------------------------")
			fmt.Println("Delete Transaction's Data")
			fmt.Println("-----------------------------------------------")
			tools.Input("", &init)
			tools.Input("insert Transaction Code to Delete: ", &transactionCode)
			Execute(conf).DeleteDataTransaction(transactionCode)
		case "3":
			fmt.Println("-----------------------------------------------")
			fmt.Println("Show All Transaction's Data")
			fmt.Println("-----------------------------------------------")
			Execute(conf).GetDataTransaction()
		}
	case "1.3":
		switch option {
		case "1":
			fmt.Println("-----------------------------------------------")
			fmt.Println("Show Daily Report's Data")
			fmt.Println("-----------------------------------------------")
			Execute(conf).GetDailyReport()
		case "2":
			fmt.Println("-----------------------------------------------")
			fmt.Println("Show Monthly Report's Data")
			fmt.Println("-----------------------------------------------")
			Execute(conf).GetMonthReport()
		case "3":
			fmt.Println("-----------------------------------------------")
			fmt.Println("Show All Reports's Data")
			fmt.Println("-----------------------------------------------")
			Execute(conf).GetAllReport()
		}
	}
}
