package models

import (
	"database/sql"
	"log"
)

//StartConnection : start connection to database
func StartConnection(engine, data string) (*sql.DB, error) {
	db, err := sql.Open(engine, data)
	if err != nil {
		log.Fatal(err)

	}
	//defer db.Close()
	pingDB(db)
	return db, nil
}

//ErrorCheck : checking connection to the database
func errorCheck(err error) {
	if err != nil {
		//panic(err.Error())
		log.Fatal(err.Error())
		//fmt.Println(err.Error())

	}
	//fmt.Print("Sucess DB Connected")

}

//PingDB : ping database
func pingDB(db *sql.DB) {
	err := db.Ping()
	errorCheck(err)

}
