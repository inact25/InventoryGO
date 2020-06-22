package main

import (
	"database/sql"
	"fmt"
	"log"
	"viper/models"
	"viper/services"
	"viper/utils"
	view "viper/views"

	_ "github.com/go-sql-driver/mysql"
)

type Env struct {
	db *sql.DB
}

//DbEngine : database engine ex:mysql
var (
	DbEngine     string
	DbUser       string
	DbPass       string
	DbHost       string
	DbPort       string
	DbConnection string
	DbSchema     string
)

func main() {
	DbEngine = utils.GetCustomConf("DbEngine", "mysql")
	DbUser = utils.GetCustomConf("DbUser", "root")
	DbPass = utils.GetCustomConf("DbPass", "password")
	DbHost = utils.GetCustomConf("DbHost", "localhost")
	DbPort = utils.GetCustomConf("DbPort", "8080")
	DbConnection = utils.GetCustomConf("DbConnection", "@tcp")
	DbSchema = utils.GetCustomConf("DbSchema", "mydatabase")

	dataSource := fmt.Sprintf("%s:%s%s(%s:%s)/%s", DbUser, DbPass, DbConnection, DbHost, DbPort, DbSchema)
	view.MainMenu()

	// ---
	mydb, err := models.StartConnection(DbEngine, dataSource)
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db: mydb}
	productService := services.ViewAllProductService(env.db)
	products := productService.GetDataProduct(1, 50)
	// //------------------------------------
	transactionService := services.ViewAllTransactionService(env.db)
	transaction := transactionService.GetDataTransaction(1, 50)

	view.ViewAllTransaction(transaction)
	view.ViewAllProduct(products)

}
