package view

import (
	"database/sql"
	"log"
	"viper/config"
	"viper/models"
)

type Env struct {
	db *sql.DB
}

func Execute(c *config.Conf) *Env {
	db, err := models.InitDB(c)
	if err != nil {
		log.Panic(err)
	}
	return &Env{db: db}
}
