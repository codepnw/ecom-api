package main

import (
	"database/sql"
	"log"

	"github.com/codepnw/ecom-api/cmd/api"
	"github.com/codepnw/ecom-api/config"
	"github.com/codepnw/ecom-api/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQL(mysql.Config{
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

	initDB(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initDB(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: connected!")
}
