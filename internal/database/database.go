package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/adonism2k/golang-hactiv8/internal/initializers"
	"github.com/adonism2k/golang-hactiv8/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
}

var count uint8

func Connect(config initializers.Config) Config {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta connect_timeout=5",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	for {
		conn, err := Open(dsn)
		if err != nil {
			log.Println("Failed to connect to Postgres. Retrying in 4 sec...")
			time.Sleep(4 * time.Second)
			count++
		} else {
			log.Println("Connected to Postgres!")

			err = conn.AutoMigrate(&model.User{}, &model.Photo{}, &model.Comment{}, &model.SocialMedia{})
			if err != nil {
				log.Println("Failed to migrate the Models.")
				log.Fatal(err)
			} else {
				log.Println("Migrated the database!")
			}

			return Config{DB: conn}
		}

		if count > 10 {
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
