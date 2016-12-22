package data

import (
	"database/sql"
	"fmt"
	"log"
)

//Database Struct for accessing the database throughout the project
var Database *sql.DB

//DialDb Opens the connection to the database
// func DialDb() error {
// 	var err error
//
// 	connString := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=%s",
// 		common.AppConfig.DatabaseServer, common.AppConfig.DatabaseUser, common.AppConfig.DatabasePassword,
// 		common.AppConfig.DatabasePort, common.AppConfig.Database, common.AppConfig.SSLMode)
//
// 	log.Println("Connecting to database:", common.AppConfig.Database)
// 	Database, err = sql.Open("postgres", connString)
//
// 	return err
// }

func DialDb() error {
	var err error

	connString := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=%s",
		"common.AppConfig.DatabaseServer", "common.AppConfig.DatabaseUser", "common.AppConfig.DatabasePassword",
		123, "common.AppConfig.Database", "common.AppConfig.SSLMode")

	log.Println("Connecting to database:", "123")
	Database, err = sql.Open("postgres", connString)

	return err
}
