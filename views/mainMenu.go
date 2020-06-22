package view

import "fmt"

func MainMenu() {
	var choose float64
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
	chooseOption(1, choose)
}

func chooseOption(id, option float64) {
	var choose float64
	switch id {
	case 1:
		switch option {
		case 1:
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
			chooseOption(1.1, choose)
		case 2:
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
			chooseOption(1.2, choose)
		case 3:
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
			chooseOption(1.3, choose)
		}

	case 1.1:
		switch option {
		case 1:
			fmt.Println("-----------------------------------------------")
			fmt.Println("Add Product's Data")
			fmt.Println("-----------------------------------------------")
		case 2:
			fmt.Println("-----------------------------------------------")
			fmt.Println("Delete Product's Data")
			fmt.Println("-----------------------------------------------")
		case 3:
			fmt.Println("-----------------------------------------------")
			fmt.Println("Show All Product's Data")
			fmt.Println("-----------------------------------------------")
		}
	case 1.2:
		switch option {
		case 1:
			fmt.Println("-----------------------------------------------")
			fmt.Println("Add Transaction's Data")
			fmt.Println("-----------------------------------------------")
		case 2:
			fmt.Println("-----------------------------------------------")
			fmt.Println("Delete Transaction's Data")
			fmt.Println("-----------------------------------------------")
		case 3:
			fmt.Println("-----------------------------------------------")
			fmt.Println("Show All Transaction's Data")
			fmt.Println("-----------------------------------------------")
		}
	case 1.3:
		switch option {
		case 1:
			fmt.Println("-----------------------------------------------")
			fmt.Println("Show Daily Report's Data")
			fmt.Println("-----------------------------------------------")
		case 2:
			fmt.Println("-----------------------------------------------")
			fmt.Println("Show Monthly Report's Data")
			fmt.Println("-----------------------------------------------")
		case 3:
			fmt.Println("-----------------------------------------------")
			fmt.Println("Show All Reports's Data")
			fmt.Println("-----------------------------------------------")
		}
	}
}
