package main

import (
	"database/sql"
	"log"

	"github.com/ayinde1993/ecom/cmd/api"
	"github.com/ayinde1993/ecom/config"
	"github.com/ayinde1993/ecom/db"
	"github.com/go-sql-driver/mysql"
)

// db connection parameter
func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	iniStorage(db)

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func iniStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
