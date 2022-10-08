package database

import (
	"assignment-2/internal/model"
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
}

var count uint8

func Connect() Config {
	dsn := os.Getenv("DSN")

	for {
		conn, err := Open(dsn)
		if err != nil {
			log.Println("Failed to connect to Postgres. Retrying in 2 sec...")
			time.Sleep(2 * time.Second)
			count++
		} else {
			log.Println("Connected to Postgres!")

			err = conn.AutoMigrate(&model.Order{}, &model.Item{})
			if err != nil {
				log.Println("Failed to migrate the Models.")
				log.Fatal(err)
			} else {
				log.Println("Migrated the database!")
			}

			return Config{DB: conn}
		}

		if count > 5 {
			log.Println("Failed to connect to Postgres. Exiting...")
			log.Println(err)
			return Config{}
		}

		continue
	}
}

func Open(dsn string) (*gorm.DB, error) {
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
