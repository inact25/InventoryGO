package config

import (
	"viper/utils"
)

type dbConf struct {
	DbUser   string
	DbPass   string
	DbHost   string
	DbPort   string
	DbSchema string
}

type Conf struct {
	Db dbConf
}

func NewAppConfig() *Conf {
	return &Conf{dbConf{
		DbUser:   utils.GetCustomConf("DbUser", "root"),
		DbPass:   utils.GetCustomConf("DbPass", "password"),
		DbHost:   utils.GetCustomConf("DbHost", "localhost"),
		DbPort:   utils.GetCustomConf("DbPort", "3306"),
		DbSchema: utils.GetCustomConf("DbSchema", "schema"),
	}}
}
