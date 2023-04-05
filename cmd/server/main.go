package main

import (
	"database/sql"
	"fmt"

	"github.com/fabio-mattos/ContabilGoBack/configs"
	_ "github.com/lib/pq"
)

func main() {

	configs, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Accessing %s ... ", configs.DBName)
	var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", configs.DBHost, configs.DBPort, configs.DBUser, configs.DBPassword, configs.DBName)

	db, err := sql.Open(configs.DBDriver, DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	defer db.Close()

}
