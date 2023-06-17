package main

import (
	"context"
	"log"
	"os"
	"plms-backend/ent"
	"plms-backend/env"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	env.LoadEnv();

	client, err := ent.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}