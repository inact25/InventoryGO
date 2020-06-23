package view

import (
	"fmt"
	"viper/config"
	"viper/tools"
)

var conf = config.NewAppConfig()

func MainMenu() {
	tools.Cls()
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
	tools.Cls()
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
		ipid:
			tools.Input("Insert Product ID  ex[P0015] (maxchar 5): ", &productCode)
			if len(productCode) > 5 {
				goto ipid
			}
			tools.Input("Insert Product Name : ", &productDesc)
		iban:
			tools.Input("Insert Brand Code  ex[B0001] (maxchar 5): ", &brandCode)
			if len(brandCode) > 5 {
				goto iban
			}
		icat:
			tools.Input("Insert Category ID  ex[C0005] (maxchar 5): ", &categoryID)
			if len(categoryID) > 5 {
				goto icat
			}
			tools.Input("Insert Product Price : ", &productPrice)
		idis:
			tools.Input("Insert Discount ID  ex[D0000] (maxchar 5): ", &discountID)
			if len(discountID) > 5 {
				goto idis
			}
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
		tid:
			tools.Input("Insert Transaction ID ex[T0005] (maxchar 5): ", &transactionCode)
			if len(transactionCode) > 5 {
				goto tid
			}
		pid:
			tools.Input("Insert Product ID : ex[P0006] (maxchar 5)", &productId)
			if len(productId) > 5 {
				goto pid
			}
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
