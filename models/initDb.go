package models

import (
	"database/sql"
	"fmt"
	"log"

	"viper/config"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB(c *config.Conf) (*sql.DB, error) {

	cfg := &mysql.Config{
		User:   c.Db.DbUser,
		Passwd: c.Db.DbPass,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%v:%v", c.Db.DbHost, c.Db.DbPort),
		DBName: c.Db.DbSchema,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Panic(err)
	}

	//Ping = check database availability
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	return db, nil
}
