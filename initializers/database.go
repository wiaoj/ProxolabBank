package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	LoadEnvVariables()
}

// var HOST = "localhost"
// var USER = "postgres"
// var PASSWORD = "postgres"
// var DATABASE_NAME = "bankApplication"
// var PORT = "5432"
// var SSL_MODE = "disable"
var POSTGRES_CONNECTION_STRING = "host=localhost user=postgres password=postgres dbname=bankapplication port=5432 sslmode=disable"

func getPostgresDsn() string {
	var HOST = os.Getenv("POSTGRES_HOST")
	var USER = os.Getenv("POSTGRES_USER")
	var PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	var DATABASE_NAME = os.Getenv("POSTGRES_DATABASE_NAME")
	var PORT = os.Getenv("POSGRES_PORT")
	var SSL_MODE = os.Getenv("SSL_MODE")
	//return "host=" + HOST + " user=" + USER + " password=" + PASSWORD + " dbname=" + DATABASE_NAME + " port=" + PORT + " sslmode=" + SSL_MODE
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", HOST, USER, PASSWORD, DATABASE_NAME, PORT, SSL_MODE)
}

func ConnectToDB() {
	var err error
	dsn := getPostgresDsn()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}

// func loadPostgresMigration() {
// 	driver, err := postgres.WithInstance(DB, &postgres.Config{})
// 	m, err := migrate.NewWithDatabaseInstance(
// 		"file:///migrations",
// 		"postgres", driver)
// 	m.Up()
// }
