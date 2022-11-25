package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Open() (*sql.DB, error) {
	viper.ReadInConfig()

	var (
		host     = "localhost"
		port     = 5432
		user     = viper.Get("DBUSER")
		password = viper.Get("DBPASSWORD")
		dbname   = "foodwars"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	return sql.Open("postgres", psqlInfo)

}
