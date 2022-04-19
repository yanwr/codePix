package db

import (
	"codePix/domain/model"
	Env "codePix/env"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "gorm.io/driver/sqlite"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func ConnectDB(env string) *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	if env != "test" {
		dsn = os.Getenv(Env.DNS)
		db, err = gorm.Open(os.Getenv(Env.DB_TYPE), dsn)
	} else {
		dsn = os.Getenv(Env.DNS_TEST)
		db, err = gorm.Open(os.Getenv(Env.DB_TYPE_TEST), dsn)
	}

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv(Env.DEBUG) == "true" {
		db.LogMode(true)
	}

	if os.Getenv(Env.AUTO_MIGRATE) == "true" {
		db.AutoMigrate(&model.Bank{}, &model.Account{}, &model.PixKey{}, &model.Transaction{})
	}

	return db
}
