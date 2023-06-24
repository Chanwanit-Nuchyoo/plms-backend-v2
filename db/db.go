package db

import (
	"log"
	"os"
	"sync"

	"plms-backend/ent"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbClient *ent.Client
	once     sync.Once
)

// GetDBClient returns the singleton instance of the ent.Client.
func GetDBClient() (*ent.Client, error) {
	var err error
	once.Do(func() {
		dbClient, err = createDBClient()
	})
	return dbClient, err
}

func createDBClient() (*ent.Client, error) {
	dbURL := os.Getenv("DB_URL")
	client, err := ent.Open("mysql", dbURL)
	if err != nil {
		log.Fatalf("failed opening connection to MySQL: %v", err)
	}

	return client, nil
}
