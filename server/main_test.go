package server

import (
	"os"
	"testing"

	"github.com/hamza72x/go-gin-gorm/accounts"
	"github.com/hamza72x/go-gin-gorm/util"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var testServer *Server

func TestMain(m *testing.M) {
	testServer = getTestServer()
	os.Exit(m.Run())
}

// panic because this is a test
func getTestServer() *Server {
	db := getTestDB()

	if err := godotenv.Load("../dev.env"); err != nil {
		panic("failed to load .env: " + err.Error())
	}

	s := New(db, os.Getenv("ENVIRONMENT"))

	return s
}

// panic because this is a test
func getTestDB() *gorm.DB {
	// rewrite test.db
	if err := util.RemoveFileIfExists("../test.db"); err != nil {
		panic("failed to remove test.db: " + err.Error())
	}

	// create test.db
	db, err := gorm.Open(sqlite.Open("../test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to create test.db: " + err.Error())
	}

	// migrations
	if err := db.AutoMigrate(&accounts.Account{}); err != nil {
		panic("failed to migrate test.db: " + err.Error())
	}

	return db
}
