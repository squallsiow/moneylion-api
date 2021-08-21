package env

import (
	"github.com/caarlos0/env/v6"
)

// Config :
var Config = struct {
	Mongo_DB_Name       string `env:"MONGO_DB_NAME" envDefault:"moneylion"`
	Mongo_DB_Connection string `env:"MONGO_DB_CONNECTION" envDefault:"mongodb+srv://admin:admin@cluster-moneylion.l6js4.mongodb.net/test"`
}{}

// Init :
func Init() {
	if err := env.Parse(&Config); err != nil {
		panic(err)
	}
}
