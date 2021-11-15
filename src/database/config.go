package databases

import (
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Connection *gorm.DB
}

var database *Database
var lock sync.Mutex

func (conn Database) SingletonInstance() *Database {
	databaseUrl := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  databaseUrl, // data source name, refer https://github.com/jackc/pgx
			PreferSimpleProtocol: true,        // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
		}), &gorm.Config{},
	)

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	log.Println("Database connected")
	conn.Connection = db

	return &conn
}

func GetDatabase() *Database {
	if database == nil {
		println("Connecting database...\n")
		database = &Database{}

		database.SingletonInstance()

	} else {
		println("Database already Connection...\n")
	}

	return database
}

func Conn() *gorm.DB {
	databaseUrl := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  databaseUrl, // data source name, refer https://github.com/jackc/pgx
			PreferSimpleProtocol: true,        // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
		}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	log.Println("Database connected")

	return db
}
