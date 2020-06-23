package main

import (
	view "viper/views"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	view.MainMenu()
}
