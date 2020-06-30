package view

import (
	"fmt"
	"viper/config"
	"viper/tools"
)

var conf = config.NewAppConfig()
var confirmation string

func Login() {
	tools.Cls()
	var userName, userPass string
	fmt.Println("-----------------------------------------------")
	fmt.Println("Login")
	fmt.Println("-----------------------------------------------")
	fmt.Printf("Insert Username : ")
	fmt.Scan(&userName)
	fmt.Printf("Insert Password : ")
	fmt.Scan(&userPass)
	Execute(conf).GetUser(userName, userPass)
}

func MainMenu() {
	tools.Cls()
	var choose string
	menu := []string{"Product's Data", "Product's Transaction", "Product's Report", "Register"}
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
	var init string
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
		case "4":
			menu := []string{"Add User", "Delete User", "Show All User"}
			fmt.Println("-----------------------------------------------")
			fmt.Println("User Master")
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
			chooseOption("1.4", choose)

		}

	case "1.1":
		switch option {
		case "1":
			var opt, productCode, productDesc, categoryID, productPrice, discountID, brandCode string
			fmt.Println("-----------------------------------------------")
			fmt.Println("Add Product's Data")
			fmt.Println("-----------------------------------------------")
			fmt.Println("Before Continue Inputing Data, please consider to view this data format")
			fmt.Println("1. view Brand Format")
			fmt.Println("2. view Category Format")
			fmt.Println("3. view Discount Format")
			fmt.Println("0. Continue Input Data")
			fmt.Printf("Make a choice > ")
			fmt.Scan(&opt)
			switch opt {
			case "1":
				fmt.Println("-----------------------------------------------")
				fmt.Println("Brand Format")
				fmt.Println("-----------------------------------------------")
				Execute(conf).GetBrand()
				fmt.Printf("wanna go back to menu ? > ")
				tools.Input("", &init)
				tools.Input("", &confirmation)
				chooseOption("1.1", "1")
			case "2":
				fmt.Println("-----------------------------------------------")
				fmt.Println("Category Format")
				fmt.Println("-----------------------------------------------")
				Execute(conf).GetCategory()
				fmt.Printf("wanna go back to menu ? > ")
				tools.Input("", &init)
				tools.Input("", &confirmation)
				chooseOption("1.1", "1")
			case "3":
				fmt.Println("-----------------------------------------------")
				fmt.Println("Discount Format")
				fmt.Println("-----------------------------------------------")
				Execute(conf).GetDiscount()
				fmt.Printf("wanna go back to menu ? > ")
				tools.Input("", &init)
				tools.Input("", &confirmation)
				chooseOption("1.1", "1")
			default:
				tools.Input("", &init)
				tools.Input("Insert ProductID  (maxchar 5): ", &productCode)
				tools.Input("Insert Product Name : ", &productDesc)
				tools.Input("Insert BrandCode  (maxchar 5): ", &brandCode)
				tools.Input("Insert CategoryID  (maxchar 5): ", &categoryID)
				tools.Input("Insert Product Price : ", &productPrice)
				tools.Input("Insert DiscountID (maxchar 5): ", &discountID)
				tools.Input("Wanna save this data ? (y/n) > ", &confirmation)
				if confirmation == "y" || confirmation == "Y" {
					Execute(conf).AddNewProduct(productCode, productDesc, categoryID, productPrice, discountID, brandCode)
					fmt.Println("Data saved, wanna go back to menu ?")
				} else {
					fmt.Println("Data not saved, wanna go back to menu ?")
				}
				fmt.Scanln()
				MainMenu()

			}

		case "2":
			var productCode string
			fmt.Println("-----------------------------------------------")
			fmt.Println("Delete Product's Data")
			fmt.Println("-----------------------------------------------")
			tools.Input("", &init)
			tools.Input("Insert Product ID to Delete: ", &productCode)
			tools.Input("Are you Serious ? (y/n) > ", &confirmation)
			if confirmation == "y" || confirmation == "Y" {
				Execute(conf).DeleteDataProduct(productCode)
				fmt.Println("Student Dropped , wanna go back to menu ?")
			} else {
				fmt.Println("Nothing happen, wanna go back to menu ?")
			}
			fmt.Scanln()
			MainMenu()
		case "3":

			fmt.Println("-----------------------------------------------")
			fmt.Println("Show All Product's Data")
			fmt.Println("-----------------------------------------------")
			Execute(conf).GetDataProduct()
			fmt.Printf("wanna go back to menu ? > ")
			tools.Input("", &init)
			tools.Input("", &confirmation)
			MainMenu()
		}
	case "1.2":
		switch option {
		case "1":
			var opt, transactionCode, productId, transactionQTY string
			fmt.Println("-----------------------------------------------")
			fmt.Println("Add Transaction's Data")
			fmt.Println("-----------------------------------------------")
			fmt.Println("Before Continue Inputing Data, please consider to view this data format")
			fmt.Println("1. view Brand Format")
			fmt.Println("0. Continue Input Data")
			fmt.Printf("Make a choice > ")
			fmt.Scan(&opt)
			switch opt {
			case "1":
				fmt.Println("-----------------------------------------------")
				fmt.Println("Brand Format")
				fmt.Println("-----------------------------------------------")
				Execute(conf).GetBrand()
				fmt.Printf("wanna go back to menu ? > ")
				tools.Input("", &init)
				tools.Input("", &confirmation)
				chooseOption("1.2", "1")
			default:
				tools.Input("", &init)
				tools.Input("Insert TransactionID (maxchar 5) : ", &transactionCode)
				tools.Input("Insert Product ID (maxchar 5) : ", &productId)
				tools.Input("Insert Quantity : ", &transactionQTY)
				tools.Input("Wanna save this data ? (y/n) > ", &confirmation)
				if confirmation == "y" || confirmation == "Y" {
					Execute(conf).AddNewTransaction(transactionCode, productId, transactionQTY)
					fmt.Println("Data saved, wanna go back to menu ?")
				} else {
					fmt.Println("Data not saved, wanna go back to menu ?")
				}
				fmt.Scanln()
				MainMenu()

			}

		case "2":
			var transactionCode string
			fmt.Println("-----------------------------------------------")
			fmt.Println("Delete Transaction's Data")
			fmt.Println("-----------------------------------------------")
			tools.Input("", &init)
			tools.Input("insert Transaction Code to Delete: ", &transactionCode)
			tools.Input("Are you Serious ? (y/n) > ", &confirmation)
			if confirmation == "y" || confirmation == "Y" {
				Execute(conf).DeleteDataTransaction(transactionCode)
				fmt.Println("Student Dropped , wanna go back to menu ?")
			} else {
				fmt.Println("Nothing happen, wanna go back to menu ?")
			}
			fmt.Scanln()
			MainMenu()

		case "3":

			fmt.Println("-----------------------------------------------")
			fmt.Println("Show All Transaction's Data")
			fmt.Println("-----------------------------------------------")
			Execute(conf).GetDataTransaction()
			fmt.Printf("wanna go back to menu ? > ")
			tools.Input("", &init)
			tools.Input("", &confirmation)
			MainMenu()
		}
	case "1.3":
		switch option {
		case "1":
			fmt.Println("-----------------------------------------------")
			fmt.Println("Show Daily Report's Data")
			fmt.Println("-----------------------------------------------")
			Execute(conf).GetDailyReport()
			fmt.Printf("wanna go back to menu ? > ")
			tools.Input("", &init)
			tools.Input("", &confirmation)
			MainMenu()

		case "2":
			fmt.Println("-----------------------------------------------")
			fmt.Println("Show Monthly Report's Data")
			fmt.Println("-----------------------------------------------")
			Execute(conf).GetMonthReport()
			fmt.Printf("wanna go back to menu ? > ")
			tools.Input("", &init)
			tools.Input("", &confirmation)
			MainMenu()

		case "3":
			fmt.Println("-----------------------------------------------")
			fmt.Println("Show All Reports's Data")
			fmt.Println("-----------------------------------------------")
			Execute(conf).GetAllReport()
			fmt.Printf("wanna go back to menu ? > ")
			tools.Input("", &init)
			tools.Input("", &confirmation)
			MainMenu()
		}
	case "1.4":
		switch option {
		case "1":
			var userName, userPass string
			fmt.Println("-----------------------------------------------")
			fmt.Println("Register User")
			fmt.Println("-----------------------------------------------")
			fmt.Printf("Insert Username : ")
			fmt.Scan(&userName)
			fmt.Printf("Insert Password :")
			fmt.Scan(&userPass)
			tools.Input("", &init)
			tools.Input("Wanna save this data ? (y/n) > ", &confirmation)
			if confirmation == "y" || confirmation == "Y" {
				Execute(conf).AddNewUser(userName, userPass)
				fmt.Println("Data saved, wanna go back to menu ?")
			} else {
				fmt.Println("Data not saved, wanna go back to menu ?")
			}
			fmt.Scanln()
			MainMenu()

		case "2":
			fmt.Println("-----------------------------------------------")
			fmt.Println("Show Monthly Report's Data")
			fmt.Println("-----------------------------------------------")
			Execute(conf).GetMonthReport()
			fmt.Printf("wanna go back to menu ? > ")
			tools.Input("", &init)
			tools.Input("", &confirmation)
			MainMenu()

		case "3":
			fmt.Println("-----------------------------------------------")
			fmt.Println("Show All User's Data")
			fmt.Println("-----------------------------------------------")
			Execute(conf).GetAllUSer()
			fmt.Printf("wanna go back to menu ? > ")
			tools.Input("", &init)
			tools.Input("", &confirmation)
			MainMenu()
		}
	}
}
