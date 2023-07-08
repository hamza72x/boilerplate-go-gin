package cmd

import (
	"errors"
	"fmt"
	"github.com/hamza72x/go-gin-gorm/accounts"
	"github.com/hamza72x/go-gin-gorm/server"
	"github.com/hamza72x/go-gin-gorm/util"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run() {
	env := os.Getenv("ENVIRONMENT")

	if env == "" {
		err := godotenv.Load("dev.env")
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				panic("could not load .env file: " + err.Error())
			}
		}
		os.Setenv("ENVIRONMENT", server.ENV_DEV)
		env = server.ENV_DEV
	}

	// envs
	if err := util.EnsureEnvs(
		"DB_HOST",
		"DB_PORT",
		"DB_NAME",
		"DB_USER",
		"DB_PASS",
		"DB_SSL",
	); err != nil {
		panic(err)
	}

	// envs post processing
	dbSSL, _ := strconv.ParseBool(os.Getenv("DB_SSL"))
	if dbSSL {
		os.Setenv("DB_SSL", "enable")
	} else {
		os.Setenv("DB_SSL", "disable")
	}

	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		panic(err)
	}

	// create db connection
	db, err := getPostgresDB()
	if err != nil {
		panic(err)
	}

	// migrations
	if err := db.AutoMigrate(&accounts.Account{}); err != nil {
		panic(err)
	}

	// server
	server := server.New(db, env)

	// run the server
	server.Run(serverPort)
}

func getPostgresDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Dhaka",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL"),
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// func getMysqlDB() (*gorm.DB, error) {
// 	dsn := fmt.Sprintf(
// 		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
// 		os.Getenv("DB_USER"),
// 		os.Getenv("DB_PASS"),
// 		os.Getenv("DB_HOST"),
// 		os.Getenv("DB_PORT"),
// 		os.Getenv("DB_NAME"),
// 		"Asia%2FDhaka",
// 	)

// 	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
// }
